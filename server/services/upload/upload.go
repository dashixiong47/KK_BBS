package upload

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/gin-gonic/gin"
	"log"
	"mime"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// FileInfoResponse 定义了文件响应的结构
type FileInfoResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Type string `json:"type"`
}

// ProcessFile 处理文件上传，并存储文件信息
func ProcessFile(c *gin.Context, size int64) (FileInfoResponse, error) {
	md5 := c.DefaultPostForm("md5", "")
	query, ok := fileQuery(md5).(FileInfoResponse)
	if ok && query.ID != 0 {
		return query, nil
	}

	file, err := c.FormFile("file")
	if err != nil {
		return FileInfoResponse{}, errors.New("get_file_error")
	}
	//if file.Size > size {
	//	return FileInfoResponse{}, errors.New("file_size_exceeds_limit")
	//}
	fileReader, err := file.Open()
	if err != nil {
		return FileInfoResponse{}, errors.New("failed_to_open_file")
	}
	defer fileReader.Close()

	md5Hash := utils.StreamToMD5(fileReader)
	//query, ok = fileQuery(md5Hash).(FileInfoResponse)
	//if ok && query.ID != 0 {
	//	return query, nil
	//}

	fileInfo, err := getFileInfo(c, file, md5Hash)
	if err != nil {
		return FileInfoResponse{}, errors.New("save_file_info_error")
	}
	save, err := warehousing(md5Hash, file.Filename, fileInfo.Type, fileInfo.Url)
	if err != nil {
		return FileInfoResponse{}, errors.New("save_file_info_error")
	}

	return FileInfoResponse{
		ID:   save.ID,
		Name: save.FileName,
		Url:  fileInfo.Url,
	}, nil
}

// fileQuery 通过md5查询文件是否已存在
func fileQuery(md5 string) any {
	if md5 == "" {
		return FileInfoResponse{}
	}
	var fileInfo models.File
	ctx := context.Background()
	if result, err := db.Rdb.Get(ctx, fmt.Sprintf("file:%v", md5)).Result(); err == nil {
		var resp FileInfoResponse
		if err := json.Unmarshal([]byte(result), &resp); err != nil {
			log.Printf("解析Redis数据错误: %v", err)
			return FileInfoResponse{}
		}
		return resp
	}
	if err := db.DB.Where("md5 = ?", md5).First(&fileInfo).Error; err != nil {
		log.Printf("查询数据库错误: %v", err)
		return FileInfoResponse{}
	}
	return FileInfoResponse{
		ID:   fileInfo.ID,
		Name: fileInfo.FileName,
		Url:  fileInfo.Url,
	}
}

// warehousing 将文件信息存储到数据库
func warehousing(md5, filename, _type, path string) (models.File, error) {
	fileInfo := models.File{
		Md5:      md5,
		FileName: filename,
		Type:     _type,
		Url:      path,
	}

	err := db.DB.Create(&fileInfo).Error
	if err != nil {
		klog.Error("创建数据库记录错误: " + err.Error())
		return models.File{}, err
	}
	// saveToRedis 将文件信息保存到Redis
	const ExpireTime = 2 * time.Hour
	data, _ := json.Marshal(FileInfoResponse{
		ID:   fileInfo.ID,
		Name: fileInfo.FileName,
		Url:  fileInfo.Url,
	})
	ctx := context.Background()
	err = db.Rdb.Set(ctx, fmt.Sprintf("file:%v", fileInfo.Md5), data, ExpireTime).Err()
	if err != nil {
		klog.Error("保存文件信息到Redis错误: " + err.Error())
	}
	return fileInfo, nil

}

// getFileInfo 根据上传的文件和其MD5获取文件信息
func getFileInfo(c *gin.Context, file *multipart.FileHeader, md5Hash string) (FileInfoResponse, error) {
	ext := filepath.Ext(file.Filename)
	mimeType := mime.TypeByExtension(ext)
	newName := fmt.Sprintf("%s%s", time.Now().Format("2006/01/02/"), md5Hash+ext)

	switch mimeType {
	case "image/jpeg", "image/png":
		url := fmt.Sprintf("files/images/%s", newName)
		err := c.SaveUploadedFile(file, url)
		if err != nil {
			return FileInfoResponse{}, errors.New("save_file_info_error")
		}
		return FileInfoResponse{Type: "image", Url: "/" + url}, nil
	case "video/mp4":
		path := fmt.Sprintf("files/videos/video/%s", newName)
		url := fmt.Sprintf("files/videos/converted/%s_hls/", newName)
		err := c.SaveUploadedFile(file, path)
		if err != nil {
			return FileInfoResponse{}, errors.New("save_file_info_error")
		}
		err = os.MkdirAll(url, 0750) // 直接创建 outputDir
		if err != nil {
			klog.Error(err.Error())
			return FileInfoResponse{}, errors.New("output_dir_create_error")
		}
		err = cutCover(path, url)
		if err != nil {
			return FileInfoResponse{}, errors.New("video_convert_error")
		}
		PushVideo(Video{
			Path: path,
			Url:  url,
		})
		//videoUrl, err := processingVideo(path, url)
		//if err != nil {
		//	return FileInfoResponse{}, errors.New("video_convert_error")
		//}
		return FileInfoResponse{Type: "video", Url: "/" + url + "index.m3u8"}, nil
	case "text/plain":
		url := fmt.Sprintf("files/videos/text/%s", newName)
		err := c.SaveUploadedFile(file, url)
		if err != nil {
			return FileInfoResponse{}, errors.New("save_file_info_error")
		}
		return FileInfoResponse{Type: "text", Url: "/" + url}, nil
	default:
		url := fmt.Sprintf("files/files/%s", newName)
		err := c.SaveUploadedFile(file, url)
		if err != nil {
			return FileInfoResponse{}, errors.New("save_file_info_error")
		}
		return FileInfoResponse{Type: "file", Url: "/" + url}, nil
	}
}

// ProcessingVideo 处理视频文件
func ProcessingVideo(path, outputDir string) error {
	output := filepath.Join(outputDir, "index.m3u8")
	//fontPath := "static/SourceHanSansCN-Bold.ttf"
	//text := "kkBBS.com"
	//cmdArgs := []string{"-i", path, "-vf", fmt.Sprintf("drawtext=text='%s':fontfile=%s:fontsize=24:fontcolor=white:x=(main_w-text_w-30):y=30", text, fontPath), "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", output}
	cmdArgs := []string{"-i", path, "-codec", "copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", output}

	cmd := exec.Command("ffmpeg", cmdArgs...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		klog.Error("FFmpeg 错误: ", stderr.String())
		return errors.New("视频转换错误")
	}
	return nil
}

// 截取封面
func cutCover(path, outputDir string) error {
	// 截取封面
	coverOutput := filepath.Join(outputDir, "cover.jpg")
	cmdArgs := []string{"-y", "-i", path, "-ss", "00:00:01", "-vframes", "1", coverOutput}
	cmd := exec.Command("ffmpeg", cmdArgs...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		klog.Error(stderr.String())
		return errors.New("封面截取错误")
	}

	return nil
}
func init() {

}

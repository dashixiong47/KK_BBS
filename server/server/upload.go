package server

import (
	"github.com/dashixiong47/KK_BBS/services/upload"
	"github.com/gin-gonic/gin"
)

type UploadServer struct{}

// var sourceUrl = "files/videos/video"
// var convertedUrl = "files/videos/converted/"
//
//	func init() {
//		services.EnsureDir(sourceUrl)
//		services.EnsureDir(convertedUrl)
//	}
func (u *UploadServer) UploadFile(c *gin.Context) (any, error) {
	size := int64(1024 * 1024 * 10)
	return upload.ProcessFile(c, size)
}

//
//// UploadVideo 获取视频
//func (u *UploadServer) UploadVideo(ctx *gin.Context) (any, error) {
//	file, err := ctx.FormFile("video")
//	if err != nil {
//		return nil, errors.New("video_not_found")
//	}
//	//判断文件大小 超出限制
//	if file.Size > 1024*1024*10 {
//		return nil, errors.New("file_size_exceeds_limit")
//	}
//	// 保存上传的视频文件
//	path := filepath.Join(sourceUrl, file.Filename)
//	err = ctx.SaveUploadedFile(file, path)
//	if err != nil {
//		klog.Error(err.Error())
//		return nil, errors.New("video_save_error")
//	}
//
//	// 转换视频为 HLS 流
//	outputDir := filepath.Join(convertedUrl, file.Filename+"_hls")
//	err = os.MkdirAll(outputDir, os.ModePerm) // 创建输出目录
//	if err != nil {
//		klog.Error(err.Error())
//		return nil, errors.New("output_dir_create_error")
//	}
//
//	output := filepath.Join(outputDir, "index.m3u8")
//	// 添加图片水印
//	// cmdArgs := []string{"-i", path, "-i", "watermark.png", "-filter_complex", "overlay=10:10", "-codec", "copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", output}
//
//	// 添加文字水印
//	fontPath := "static/SourceHanSansCN-Bold.ttf" // 替换为您系统上的默认字体路径
//	text := "kkBBS.com"                           // 水印的文字内容
//	cmdArgs := []string{"-i", path, "-vf", fmt.Sprintf("drawtext=text='%s':fontfile=%s:fontsize=24:fontcolor=white:x=(main_w-text_w-30):y=30", text, fontPath), "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", output}
//
//	// 正常转为 HLS 流
//	// cmdArgs := []string{"-i", path, "-codec", "copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", output}
//	cmd := exec.Command("ffmpeg", cmdArgs...)
//
//	// 打印 FFmpeg 命令
//	// log.Println("Executing FFmpeg command:", "ffmpeg", strings.Join(cmdArgs, " "))
//	var stderr bytes.Buffer
//	cmd.Stderr = &stderr
//
//	err = cmd.Run()
//	if err != nil {
//		klog.Error(stderr.String()) // 打印 FFmpeg 错误输出
//		return nil, errors.New("video_convert_error")
//	}
//
//	return map[string]string{"hls_playlist": output}, nil
//}

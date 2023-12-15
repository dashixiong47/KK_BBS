package upload

import (
	"sync"

	"github.com/dashixiong47/KK_BBS/utils/klog"
)

type Video struct {
	Path string
	Url  string
}

var (
	videoPond = make(chan Video, 10)
	done      = make(chan struct{})
	wg        sync.WaitGroup
	once      sync.Once
)

func init() {
	const workerCount = 5
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go videoWorker()
	}
}

func videoWorker() {
	defer wg.Done()
	for {
		select {
		case info := <-videoPond:
			err := ProcessingVideo(info.Path, info.Url)
			if err != nil {
				klog.Error("视频处理错误: ", err.Error(), " 文件: ", info.Path)
			}
		case <-done:
			return
		}
	}
}

func PushVideo(info Video) {
	select {
	case videoPond <- info:
	case <-done:
		klog.Error("尝试推送到已关闭的通道: ", info.Path)
	}
}

func CloseVideoProcessing() {
	once.Do(func() {
		close(done)
		wg.Wait()
		close(videoPond)
	})
}

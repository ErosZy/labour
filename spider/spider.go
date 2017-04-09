package spider

import (
	"strconv"
	"strings"
	"time"

	"github.com/ErosZy/labour/filterPool"
	dl "github.com/ErosZy/labour/middleware/downloader"
	pl "github.com/ErosZy/labour/middleware/pipeliner"
	"github.com/ErosZy/labour/models"
	"github.com/ErosZy/labour/process"
	"github.com/ErosZy/labour/scheduler"
	"github.com/ErosZy/labour/urlHeap"

	"github.com/ErosZy/singoriensis"
	"github.com/ErosZy/singoriensis/interfaces"
)

func Run(taskInfo *models.TaskInfoModel) {
	taskName := strconv.Itoa(int(time.Now().Unix()))
	spider := singoriensis.NewSpider(taskName, process.NewProcess(taskInfo))

	downloader := getDownloader(taskInfo)
	scheduler := getScheduler()
	pipeliner := getPipeliner()

	spider.SetDownloader(downloader)
	spider.SetScheduler(scheduler)
	spider.SetPipeliner(pipeliner)
	spider.SetThreadNum(taskInfo.ThreadNum)
	spider.SetTimeout((time.Duration)(taskInfo.CloseTime) * time.Minute)

	targetUrls := strings.Split(taskInfo.TargetUrl, ";")
	for _, v := range targetUrls {
		spider.AddUrl(v)
	}

	spider.Run()
}

func getDownloader(taskInfo *models.TaskInfoModel) *singoriensis.Downloader {
	downloader := singoriensis.NewDownloader()
	downloader.SetRetryMaxCount(taskInfo.RetryMaxCount)
	downloader.SetSleepTime((time.Duration)(taskInfo.SleepTime) * time.Millisecond)

	requests := make([]interfaces.RequestInterface, 0, taskInfo.ThreadNum)
	for i := 0; i < taskInfo.ThreadNum; i++ {
		requests = append(requests, singoriensis.NewRequest(downloader))
	}

	downloader.SetRequests(requests)

	proxyDownloaderMiddleware := dl.NewProxyDownloaderMiddleware(taskInfo)
	downloader.RegisterMiddleware(proxyDownloaderMiddleware)

	downloaderMiddleware := dl.NewDownloaderMiddleware(taskInfo)
	downloader.RegisterMiddleware(downloaderMiddleware)

	return downloader
}

func getScheduler() *scheduler.ListScheduler {
	bfUrlHeap := urlHeap.NewBFUrlHeap()
	listScheduler := scheduler.NewListScheduler()
	listScheduler.SetUrlHeap(bfUrlHeap)
	return listScheduler
}

func getPipeliner() *singoriensis.Pipeliner {
	pipeliner := singoriensis.NewPipeliner()
	filterPool := filterPool.NewFilterPool()

	logPipelinerMiddleware := pl.NewLogPipelinerMiddleware(filterPool)
	pipeliner.RegisterMiddleware(logPipelinerMiddleware)

	return pipeliner
}

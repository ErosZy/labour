package downloader

import (
	"net/http"
	"time"

	"github.com/ErosZy/labour/models"
	"github.com/ErosZy/labour/router"

	_common "github.com/ErosZy/singoriensis/common"
	_error "github.com/ErosZy/singoriensis/error"
)

var globalHeaders *models.RequestHeaderModel

type DownloaderMiddleware struct {
	router    *router.Router
	taskInfo  *models.TaskInfoModel
	transport *http.Transport
}

func NewDownloaderMiddleware(taskInfo *models.TaskInfoModel) *DownloaderMiddleware {
	schedulerItems := taskInfo.Schedulers
	router := router.NewRouter()

	for _, v := range schedulerItems {
		schedulerHeaders := models.NewRequestHeaderModel(taskInfo.RequestHeaders)
		router.Add(v.Route, getDownloaderRouterHandler(schedulerHeaders, taskInfo.Method))
	}

	return &DownloaderMiddleware{
		router:   router,
		taskInfo: taskInfo,
		transport: &http.Transport{
			DisableCompression:    false,
			DisableKeepAlives:     false,
			ResponseHeaderTimeout: 0,
		},
	}
}

func (self *DownloaderMiddleware) SetClient(stop *bool, client *http.Client) {
	client.Timeout = (time.Duration)(self.taskInfo.RequestTimeout) * time.Second
	client.Transport = self.transport
}

func (self *DownloaderMiddleware) SetRequest(stop *bool, req *http.Request, err *_error.RequestError) {
	self.router.Match(req.URL.String(), req)
}

func (self *DownloaderMiddleware) GetResponse(stop *bool, page *_common.Page, err *_error.ResponseError) {
}

func (self *DownloaderMiddleware) Error(stop *bool, client *http.Client, err error) {
}

func getDownloaderRouterHandler(schedulerHeaders *models.RequestHeaderModel, method string) router.RouterHandlerFunc {
	return func(urlStr string, arg ...interface{}) {
		headers := schedulerHeaders.RequestHeaders
		req := arg[0].(*http.Request)
		req.Method = method

		for _, v := range headers {
			req.Header.Add(v.Key, v.Value)
		}
	}
}

package downloader

import (
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/ErosZy/labour/common"
	"github.com/ErosZy/labour/models"

	_common "github.com/ErosZy/singoriensis/common"
	_error "github.com/ErosZy/singoriensis/error"
)

type ProxyDownloaderMiddleware struct {
	proxys []models.ProxyModel
}

func NewProxyDownloaderMiddleware(taskInfo *models.TaskInfoModel) *ProxyDownloaderMiddleware {
	return &ProxyDownloaderMiddleware{taskInfo.Proxy}
}

func (self *ProxyDownloaderMiddleware) SetClient(stop *bool, client *http.Client) {
	if len(self.proxys) > 0 {
		proxy := self.getRandomProxy()
		uri, err := url.Parse("http://" + proxy.Ip + ":" + proxy.Port)

		if err != nil {
			common.Logger(common.LOG_WARNING, err.Error())
			return
		}

		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(uri),
		}
	}
}

func (self *ProxyDownloaderMiddleware) SetRequest(stop *bool, req *http.Request, err *_error.RequestError) {
}

func (self *ProxyDownloaderMiddleware) GetResponse(stop *bool, page *_common.Page, err *_error.ResponseError) {
}

func (self *ProxyDownloaderMiddleware) Error(stop *bool, client *http.Client, err error) {
	self.SetClient(stop, client)
}

func (self *ProxyDownloaderMiddleware) getRandomProxy() *models.ProxyModel {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(self.proxys))
	return &self.proxys[index]
}

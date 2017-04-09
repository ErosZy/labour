package models

import (
	"encoding/json"

	"github.com/bitly/go-simplejson"
)

type TaskInfoModel struct {
	TargetUrl      string
	ThreadNum      int
	RetryMaxCount  int
	SleepTime      int
	CloseTime      int
	RequestTimeout int
	Method         string
	RequestHeaders string
	Proxy          []ProxyModel
	Schedulers     []SchedulerItemModel
	Pages          []PageItemModel
}

func NewTaskInfoModel(config *simplejson.Json) *TaskInfoModel {
	taskInfo := TaskInfoModel{}

	taskInfo.TargetUrl, _ = config.GetPath("targetUrl").String()
	taskInfo.ThreadNum, _ = config.GetPath("threadNum").Int()
	taskInfo.RetryMaxCount, _ = config.GetPath("retryMaxCount").Int()
	taskInfo.SleepTime, _ = config.GetPath("sleepTime").Int()
	taskInfo.CloseTime, _ = config.GetPath("closeTime").Int()
	taskInfo.RequestTimeout, _ = config.GetPath("requestTimeout").Int()
	taskInfo.Method, _ = config.GetPath("method").String()

	body, _ := config.GetPath("headers").Encode()
	taskInfo.RequestHeaders = string(body)

	body, _ = config.GetPath("proxy").Encode()
	json.Unmarshal(body, &taskInfo.Proxy)

	body, _ = config.GetPath("schedulers").Encode()
	json.Unmarshal(body, &taskInfo.Schedulers)

	body, _ = config.GetPath("pages").Encode()
	json.Unmarshal(body, &taskInfo.Pages)

	return &taskInfo
}

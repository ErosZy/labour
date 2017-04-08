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
	RequestHeaders string
	Schedulers     []SchedulerItemModel
	Pages          []PageItemModel
}

func NewTaskInfoModel(config *simplejson.Json) *TaskInfoModel {
	taskInfo := TaskInfoModel{}

	taskInfo.TargetUrl, _ = config.GetPath("targetUrl").String()
	taskInfo.ThreadNum, _ = config.GetPath("threadNum").Int()
	taskInfo.RetryMaxCount, _ = config.GetPath("sleepTime").Int()
	taskInfo.CloseTime, _ = config.GetPath("closeTime").Int()
	taskInfo.RequestTimeout, _ = config.GetPath("requestTimeout").Int()

	body, _ := config.GetPath("headers").Encode()
	taskInfo.RequestHeaders = string(body)

	body, _ = config.GetPath("schedulers").Encode()
	json.Unmarshal(body, &taskInfo.Schedulers)

	body, _ = config.GetPath("pages").Encode()
	json.Unmarshal(body, &taskInfo.Pages)

	return &taskInfo
}

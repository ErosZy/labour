package main

import (
	"io/ioutil"
	"os"
	"runtime"

	"github.com/ErosZy/labour/common"
	"github.com/ErosZy/labour/models"
	"github.com/ErosZy/labour/spider"

	"github.com/bitly/go-simplejson"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		common.Logger(common.LOG_FATAL, "you should start like: labour.exe config/task.json")
	}

	config := loadConfig(args[1])
	taskInfo := models.NewTaskInfoModel(config)
	runtime.GOMAXPROCS(runtime.NumCPU())
	spider.Run(taskInfo)

	common.Logger(common.LOG_INFO, "task already complete all ^_^")
}

func loadConfig(filepath string) *simplejson.Json {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		common.Logger(common.LOG_FATAL, "read "+filepath+" faild:"+err.Error())
	}

	json, err := simplejson.NewJson(data)
	if err != nil {
		common.Logger(common.LOG_FATAL, "parse config faild:"+err.Error())
	}

	return json
}

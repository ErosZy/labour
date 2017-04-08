package main

import (
	"io/ioutil"
	"labour/common"
	"labour/models"
	"labour/spider"
	"os"
	"runtime"

	"github.com/bitly/go-simplejson"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		common.Logger(common.LOG_FATAL, "you should start like: labour.exe config/task.json")
		os.Exit(0)
	}

	config := loadConfig(args[1])
	taskInfo := models.NewTaskInfoModel(config)
	runtime.GOMAXPROCS(runtime.NumCPU())
	spider.Run(taskInfo)

	common.Logger(common.LOG_INFO, "task already complete all #^-^#")
}

func loadConfig(filepath string) *simplejson.Json {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		common.Logger(common.LOG_FATAL, "read "+filepath+" faild:"+err.Error())
		os.Exit(0)
	}

	json, err := simplejson.NewJson(data)
	if err != nil {
		common.Logger(common.LOG_FATAL, "parse config faild:"+err.Error())
		os.Exit(0)
	}

	return json
}

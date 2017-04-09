package pipeliner

import (
	"github.com/ErosZy/labour/common"
	"github.com/ErosZy/labour/filterPool"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var lock sync.Mutex

type LogPipelinerMiddleware struct {
	logger     *log.Logger
	logFile    *os.File
	filterPool *filterPool.FilterPool
}

func NewLogPipelinerMiddleware(filterPool *filterPool.FilterPool) *LogPipelinerMiddleware {
	realPath := "./" + strconv.Itoa(int(time.Now().Unix())) + ".log"
	logFile, err := os.OpenFile(realPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)

	if err != nil {
		common.Logger(common.LOG_FATAL, err.Error())
	}

	logger := log.New(logFile, "", log.Ldate|log.Ltime)

	return &LogPipelinerMiddleware{
		logger:     logger,
		logFile:    logFile,
		filterPool: filterPool,
	}
}

func (self *LogPipelinerMiddleware) GetItems(stop *bool, pipelinerItems ...interface{}) {
	lock.Lock()
	defer lock.Unlock()

	for _, v := range pipelinerItems {
		item := v.(common.PipelinerItem)
		if !self.filterPool.Contain(item.MainKeyValue) {
			self.logger.Println(item.BaseUrl + " " + item.MainKey + " " + item.Body + " " + item.BodyType)
		} else {
			*stop = true
			break
		}
	}
}

func (self *LogPipelinerMiddleware) Close() {
	self.logFile.Close()
}

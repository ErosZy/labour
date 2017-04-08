package router

import (
	"labour/common"
	"regexp"
)

type Router struct {
	patterns []*RouterItem
}

func NewRouter() *Router {
	return &Router{
		patterns: make([]*RouterItem, 0),
	}
}

func (self *Router) Add(pattern string, handler RouterHandlerFunc) {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		common.Logger(common.LOG_FATAL, "Router pattern compile error!")
	} else {
		self.patterns = append(self.patterns, NewRouterItem(reg, handler))
	}

}

func (self *Router) Match(url string, params ...interface{}) {
	patterns := self.patterns
	for _, v := range patterns {
		if v.reg.MatchString(url) {
			v.handler(url, params...)
			break
		}
	}
}

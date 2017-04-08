package router

import (
	"regexp"
)

type RouterHandlerFunc func(string, ...interface{})

type RouterItem struct {
	reg     *regexp.Regexp
	handler RouterHandlerFunc
}

func NewRouterItem(reg *regexp.Regexp, handler RouterHandlerFunc) *RouterItem {
	return &RouterItem{reg, handler}
}

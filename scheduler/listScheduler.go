package scheduler

import (
	"container/list"

	_common "github.com/ErosZy/labour/common"

	"sync"

	"github.com/ErosZy/singoriensis/common"
	"github.com/ErosZy/singoriensis/interfaces"
)

var lock sync.Mutex

type ListScheduler struct {
	list        *list.List
	urlHeap     interfaces.UrlHeapInterface
	middlewares []interfaces.SchedulerMiddlewareInterface
}

func NewListScheduler() *ListScheduler {
	return &ListScheduler{
		list:        list.New(),
		middlewares: make([]interfaces.SchedulerMiddlewareInterface, 0),
	}
}

func (self *ListScheduler) GetElemCount() int {
	return self.list.Len()
}

func (self *ListScheduler) SetUrlHeap(urlHeap interfaces.UrlHeapInterface) {
	self.urlHeap = urlHeap
}

func (self *ListScheduler) RegisterMiddleware(mw interfaces.SchedulerMiddlewareInterface) {
	self.middlewares = append(self.middlewares, mw)
}

func (self *ListScheduler) CallMiddlewareMethod(name string, params []interface{}) {
	if len(self.middlewares) != 0 {
		common.CallObjMethod(self.middlewares, name, params)
	}
}

func (self *ListScheduler) AddElementItem(elem common.ElementItem, isForce bool) {
	lock.Lock()
	defer lock.Unlock()

	self.CallMiddlewareMethod("ElementItemIn", []interface{}{&elem})

	if self.urlHeap == nil {
		_common.Logger(_common.LOG_FATAL, "scheduler's urlHeap is empty")
	}

	if isForce || (elem.UrlStr != "" && !self.urlHeap.Contain(elem)) {
		self.list.PushBack(elem)
	}
}

func (self *ListScheduler) ShiftElementItem() interface{} {
	lock.Lock()
	defer lock.Unlock()

	elem := self.list.Front()

	if elem == nil {
		return nil
	}

	self.CallMiddlewareMethod("ElementItemOut", []interface{}{&elem.Value})
	return self.list.Remove(elem)
}

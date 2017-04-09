package urlHeap

import (
	"github.com/ErosZy/singoriensis/common"

	"github.com/AndreasBriese/bbloom"
)

type BFUrlHeap struct {
	bf bbloom.Bloom
}

func NewBFUrlHeap() *BFUrlHeap {
	return &BFUrlHeap{
		bf: bbloom.New(float64(1<<16), float64(0.01)),
	}
}

func (self *BFUrlHeap) Contain(elem common.ElementItem) bool {
	url := []byte(elem.UrlStr)
	isIn := self.bf.Has(url)

	if !isIn {
		self.bf.Add(url)
	}

	return isIn
}

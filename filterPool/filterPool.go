package filterPool

type FilterPool struct {
	kmap map[string]bool
}

func NewFilterPool() *FilterPool {
	return &FilterPool{make(map[string]bool)}
}

func (self *FilterPool) Contain(key string) bool {
	if !self.kmap[key] {
		self.kmap[key] = true
		return false
	}

	return true
}

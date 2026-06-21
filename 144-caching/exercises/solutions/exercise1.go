package solutions

type LRU struct {
	cap   int
	order []string
	data  map[string]int
}

func NewLRU(cap int) *LRU {
	return &LRU{cap: cap, data: make(map[string]int)}
}

func (c *LRU) Put(key string, val int) {
	if c.cap == 0 { return }
	if _, ok := c.data[key]; ok {
		c.data[key] = val
		return
	}
	if len(c.order) >= c.cap {
		delete(c.data, c.order[0])
		c.order = c.order[1:]
	}
	c.order = append(c.order, key)
	c.data[key] = val
}

func (c *LRU) Get(key string) (int, bool) {
	v, ok := c.data[key]
	return v, ok
}
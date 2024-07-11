package member

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Gatherer struct {
	mu       sync.Mutex
	schedule map[string]*call
}

func NewGatherer() (gatherer *Gatherer) {
	gatherer = &Gatherer{
		schedule: make(map[string]*call),
	}
	return
}

func (g *Gatherer) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()

	if c, ok := g.schedule[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.schedule[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.schedule, key)
	g.mu.Unlock()

	return c.val, c.err
}

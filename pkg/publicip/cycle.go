package publicip

import (
	"sync"
)

type cycler interface {
	next() Provider
}

type cyclerImpl struct {
	sync.Mutex
	counter   int
	providers []Provider
}

func newCycler(providers []Provider) cycler {
	return &cyclerImpl{
		providers: providers,
	}
}

func (c *cyclerImpl) next() Provider {
	c.Lock()
	defer c.Unlock()
	provider := c.providers[c.counter]
	c.counter++
	if c.counter == len(c.providers) {
		c.counter = 0
	}
	return provider
}

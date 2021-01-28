package publicip

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_newCycler(t *testing.T) {
	t.Parallel()
	providers := []Provider{Opendns, Ifconfig}
	cycler := newCycler(providers)
	require.NotNil(t, cycler)
	provider := cycler.next()
	assert.Equal(t, Opendns, provider)
}

func Test_next(t *testing.T) {
	t.Parallel()
	cycler := &cyclerImpl{
		providers: []Provider{Opendns, Ifconfig},
	}
	var provider Provider
	provider = cycler.next()
	assert.Equal(t, provider, Opendns)
	provider = cycler.next()
	assert.Equal(t, provider, Ifconfig)
	provider = cycler.next()
	assert.Equal(t, provider, Opendns)
}

func Test_next_RaceCondition(t *testing.T) {
	// Run with -race flag
	t.Parallel()
	const workers = 5
	const loopSize = 101
	c := &cyclerImpl{
		providers: []Provider{Opendns, Ifconfig, Ipify, Ipinfo},
	}
	ready := make(chan struct{})
	wg := &sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			<-ready
			for i := 0; i < loopSize; i++ {
				c.next()
			}
			wg.Done()
		}()
	}
	close(ready)
	wg.Wait()
	assert.Equal(t, (workers*loopSize)%len(c.providers), c.counter)
}

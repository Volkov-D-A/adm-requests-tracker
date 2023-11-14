package closer

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

type Closer struct {
	mu    sync.Mutex
	funcs map[string]func(ctx context.Context) error
}

func NewAppCloser() *Closer {
	funcs := make(map[string]func(ctx context.Context) error)
	return &Closer{funcs: funcs}
}

func (c *Closer) Add(name string, f func(ctx context.Context) error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.funcs[name] = f
}

func (c *Closer) Remove(name string) {
	delete(c.funcs, name)
}

func (c *Closer) Close(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	msgs := make([]string, 0, len(c.funcs))
	complete := make(chan struct{}, 1)

	go func() {
		for _, f := range c.funcs {
			err := f(ctx)
			if err != nil {
				msgs = append(msgs, fmt.Sprintf("[!] %v", err))
			}
		}

		complete <- struct{}{}
	}()

	select {
	case <-complete:
		break
	case <-ctx.Done():
		return fmt.Errorf("shutdown cancelled: %v", ctx.Err())
	}

	if len(msgs) > 0 {
		return fmt.Errorf("shutdown finished with error(s): \n%s", strings.Join(msgs, "\n"))
	}
	return nil
}

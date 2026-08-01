package redis

import (
	"context"
	"sync"
)

type Pipeline struct {
	mu   sync.Mutex
	cmds []string
}

func (p *Pipeline) Exec(ctx context.Context) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if err := ctx.Err(); err != nil {
		return err
	}
	p.cmds = nil
	return nil
}

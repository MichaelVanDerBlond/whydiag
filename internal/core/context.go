package core

import (
	"context"
	"time"
)

// Context содержит контекст выполнения диагностики.
// На R0.1 он минимальный, но в следующих релизах будет
// расширяться без изменения существующего API.
type Context struct {
	Context context.Context
	Started time.Time
}

// NewContext создает новый контекст выполнения.
func NewContext() *Context {
	return &Context{
		Context: context.Background(),
		Started: time.Now(),
	}
}

// Uptime возвращает длительность выполнения.
func (c *Context) Uptime() time.Duration {
	return time.Since(c.Started)
}

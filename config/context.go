package config

import (
	"context"
	"go.uber.org/zap"
)

type Ctx struct {
	Context context.Context
	cancel  context.CancelFunc
	Cfg     *Config
	Log     *zap.Logger
}

func NewCtx(cfg *Config, logger *zap.Logger) *Ctx {
	ctx, cancel := context.WithCancel(context.Background())
	return &Ctx{
		Context: ctx,
		cancel:  cancel,
		Cfg:     cfg,
		Log:     logger,
	}
}

func (c *Ctx) Cancel() {
	c.cancel()
}

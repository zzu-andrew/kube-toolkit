package config

import (
	"context"
	"go.uber.org/zap"
)

type Ctx struct {
	Context context.Context
	Cfg     *Config
	Logger  *zap.Logger
}

func NewCtx(cfg *Config, logger *zap.Logger) *Ctx {
	return &Ctx{
		Cfg:    cfg,
		Logger: logger,
	}
}

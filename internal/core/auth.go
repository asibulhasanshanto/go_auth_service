package core

import "go.uber.org/zap"

type Auth struct {
	log *zap.Logger
}

func NewAuth(log *zap.Logger) *Auth {
	return &Auth{log: log}
}


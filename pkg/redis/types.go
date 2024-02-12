package redis

import (
	"context"

	"github.com/raanfefu/web-srv-tls-template/pkg/commons"
)

type RedisClient interface {
	DefaultConfig()
	DefaultConfigDefer()
	Get(key string)
}

type ParametersCfg struct {
	CfgFile       string
	TlsMode       bool
	AnonymousMode bool
	Addr          string
	Username      string
	Password      string
	Database      uint
	Key           string
	Certificate   string
	RootCa        string
}

type impl struct {
	TypeSettings commons.TypeSettings
	Ctx          context.Context
	Parameters   ParametersCfg
}

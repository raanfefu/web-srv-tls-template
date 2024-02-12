package redis

import (
	"context"
	"flag"

	"github.com/raanfefu/web-srv-tls-template/pkg/commons"
)

func NewRedisClient(context context.Context, vType commons.TypeSettings) RedisClient {
	obj := &impl{
		TypeSettings: vType,
		Ctx:          context,
	}
	return obj
}

func (r *impl) DefaultConfigDefer() {
	if r.TypeSettings == commons.JsonSettings {
		flag.StringVar(&r.Parameters.CfgFile, "redis-cfg", "", "ruta del archivo de configuración de redis")
	}

	if r.TypeSettings == commons.ArgumentSettings {
		flag.BoolVar(&r.Parameters.TlsMode, "redis-ssl", false, "si este flag esta presente son requeridos los -redis-key, -redis-crt y -redis-rootca")
		flag.BoolVar(&r.Parameters.AnonymousMode, "redis-anom", false, "si este flag esta presente son requeridos los -redis-user y -redis-pass")
		flag.StringVar(&r.Parameters.Addr, "redis-addr", "", "host redis")
		flag.StringVar(&r.Parameters.Username, "redis-user", "", "username redis")
		flag.StringVar(&r.Parameters.Password, "redis-pass", "", "password redis")
		flag.UintVar(&r.Parameters.Database, "redis-db", 0, "password redis")
		flag.StringVar(&r.Parameters.Key, "redis-key", "", "absoluted path file key pem")
		flag.StringVar(&r.Parameters.Certificate, "redis-crt", "", "absoluted path file cert pem")
		flag.StringVar(&r.Parameters.RootCa, "redis-rootca", "", "absoluted path file cert pem")
	}
}

func (r *impl) DefaultConfig() {
	r.DefaultConfigDefer()
	flag.Parse()
}

func (r *impl) Get(key string) {

}

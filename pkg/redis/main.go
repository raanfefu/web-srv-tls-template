package redis

import (
	"flag"
	"os"

	c "github.com/raanfefu/web-srv-tls-template/common"
)

type RedisClient interface {
	CheckServerParams()
	InitClientRedis()
	Connect()
}

type impl struct {
	Config *RedisParameters
}

type RedisParameters struct {
	ModeSSL     bool
	Addr        string
	Username    string
	Password    string
	Certificate string
	Key         string
	RootCAs     string
}

func NewClient() RedisClient {
	return &impl{
		Config: &RedisParameters{},
	}
}

func (r *impl) CheckServerParams() {
	flag.BoolVar(&r.Config.ModeSSL, "redis-ssl", false, "true / false")
	flag.StringVar(&r.Config.Addr, "redis-addr", "", "host:port address redis server ")
	flag.StringVar(&r.Config.Username, "redis-user", "", "username redis")
	flag.StringVar(&r.Config.Password, "redis-pass", "", "password redis")
	flag.StringVar(&r.Config.Key, "redis-key", "", "absoluted path file key pem")
	flag.StringVar(&r.Config.Certificate, "redis-crt", "", "absoluted path file cert pem")
	flag.StringVar(&r.Config.Certificate, "redis-rootca", "", "absoluted path file cert pem")

	if c.StringIsRequiered(&r.Config.Addr) != nil {
		flag.PrintDefaults()
		os.Exit(-1)
	}
}

func (r *impl) InitClientRedis() {
	//redis.NewClient(&r.Config)
}

func (r *impl) Connect() {
	//redis.NewClient(&r.Config)
}

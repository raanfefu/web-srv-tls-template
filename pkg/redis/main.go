package redis

/*
import (
	"context"
	"flag"
	"log"
	"os"

	c "github.com/raanfefu/web-srv-tls-template/pkg/commons"
	redis "github.com/redis/go-redis/v9"
)

type RedisClient interface {
	CheckServerParams()
	//DefaultConfig()
	InitClientRedis()
	Connect()
}

type impl struct {
	Params *RedisParameters
	Config *redis.Options
	Client *redis.Client
	Ctx    context.Context
}

type RedisParameters struct {
	Addr     string
	Database uint

	ModeSSL     bool
	Certificate string
	Key         string
	RootCAs     string

	Anom     bool
	Username string
	Password string
}

func NewClient(parse ...bool) RedisClient {
	r := &impl{
		Params: &RedisParameters{},
	}
	flag.BoolVar(&r.Params.ModeSSL, "redis-ssl", false, "si este flag esta presente son requeridos los -redis-key, -redis-crt y -redis-rootca")
	flag.BoolVar(&r.Params.Anom, "redis-anom", false, "si este flag esta presente son requeridos los -redis-user y -redis-pass")
	flag.StringVar(&r.Params.Addr, "redis-addr", "", "host redis")
	flag.StringVar(&r.Params.Username, "redis-user", "", "username redis")
	flag.StringVar(&r.Params.Password, "redis-pass", "", "password redis")
	flag.UintVar(&r.Params.Database, "redis-db", 0, "password redis")
	flag.StringVar(&r.Params.Key, "redis-key", "", "absoluted path file key pem")
	flag.StringVar(&r.Params.Certificate, "redis-crt", "", "absoluted path file cert pem")
	flag.StringVar(&r.Params.Certificate, "redis-rootca", "", "absoluted path file cert pem")
	if len(parse) > 0 && parse[0] {
		flag.Parse()
	}
	return r
}

func (r *impl) CheckServerParams() {

	if c.StringIsRequiered(&r.Params.Addr) != nil {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	if !r.Params.Anom {
		if c.StringIsRequiered(&r.Params.Username) != nil {
			flag.PrintDefaults()
			os.Exit(-1)
		}

		if c.StringIsRequiered(&r.Params.Password) != nil {
			flag.PrintDefaults()
			os.Exit(-1)
		}
	}

	if r.Params.ModeSSL {
		if c.StringIsRequiered(&r.Params.Certificate) != nil {
			flag.PrintDefaults()
			os.Exit(-1)
		}

		if c.StringIsRequiered(&r.Params.Key) != nil {
			flag.PrintDefaults()
			os.Exit(-1)
		}

		if c.StringIsRequiered(&r.Params.RootCAs) != nil {
			flag.PrintDefaults()
			os.Exit(-1)
		}

	}
	log.Printf("Reading redis parameters... Done ✓")

}

func (r *impl) InitClientRedis() {
	log.Printf("anom : %t", r.Params.Anom)
	log.Printf("ssl : %t", r.Params.ModeSSL)

	r.Config = &redis.Options{
		Addr: r.Params.Addr,
		DB:   int(r.Params.Database),
	}

	if r.Params.ModeSSL {
		log.Printf("Not Support SSL Mode")
		os.Exit(-1)
	}

	if !r.Params.Anom {
		log.Printf("AQUIIII")
		r.Config.Username = r.Params.Username
		r.Config.Password = r.Params.Password
	}

	r.Ctx = context.TODO()
	log.Println("Initializing redis server... Done ✓")
}

func (r *impl) Connect() {
	r.Client = redis.NewClient(r.Config)
	rd := r.Client.Ping(r.Ctx)
	_, e := rd.Result()
	if e != nil {
		log.Printf("Error connecting redis server... Failed \n %s  ", e.Error())
		os.Exit(-1)
	}
	log.Println("Connecting redis server... Done ✓")

}
*/

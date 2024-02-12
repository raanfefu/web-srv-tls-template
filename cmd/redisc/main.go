package main

import (
	"context"
	"log"

	"github.com/raanfefu/web-srv-tls-template/pkg/commons"
	"github.com/raanfefu/web-srv-tls-template/pkg/redis"
)

func main() {

	client := redis.NewRedisClient(context.TODO(), commons.ArgumentSettings)
	log.Println(client)
}

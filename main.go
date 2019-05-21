package main

import (
	"github.com/RXDA/RedisAPI/config"
	"github.com/RXDA/RedisAPI/router"
)

func main() {
	env := "develop"
	config.Init(env)
	conn := config.Conn.RedisPool.Get()
	defer conn.Close()

	r := router.Router()
	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

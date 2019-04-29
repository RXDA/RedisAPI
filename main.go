package main

import (
	"daichaogo/config"
	"daichaogo/service/keys"
	"daichaogo/service/sys"
	"github.com/gin-gonic/gin"
)

func main() {
	env := "develop"
	config.Init(env)
	conn := config.Conn.RedisPool.Get()
	defer conn.Close()

	r := gin.Default()
	r.GET("/keys", sys.GetAllKeys)
	r.GET("/infos", sys.GetServerInfo)
	r.GET("/type", keys.Type)
	r.Run() // listen and serve on 0.0.0.0:8080
}

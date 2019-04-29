package keys

import (
	"daichaogo/config"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
)

func Type(c *gin.Context) {
	key := c.Query("key")
	conn := config.Conn.RedisPool.Get()
	defer conn.Close()

	keyExists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		log.Fatal(err)
	}
	//key dose not exists
	if !keyExists {
		c.JSON(200, "key dose not exists")
		return
	}
	keyType, err := redis.String(conn.Do("TYPE", key))
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, struct {
		Key  string
		Type string
	}{key, keyType})
}

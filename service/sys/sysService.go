package sys

import (
	"github.com/RXDA/RedisAPI/config"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"strings"
)

//get all keys in redis
func GetAllKeys(c *gin.Context) {
	conn := config.Conn.RedisPool.Get()
	defer conn.Close()
	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, keys)
}

//get redis server info
func GetServerInfo(c *gin.Context){
	conn := config.Conn.RedisPool.Get()
	defer conn.Close()
	infos, err := redis.String(conn.Do("INFO", "server"))
	if err != nil {
		log.Fatal(err)
	}
	//format
	result :=make(map[string]string)
	infoArray := strings.Split(infos,"\r\n")
	for _,v :=range infoArray{
		infoOne := strings.Split(v,":")
		if len(infoOne)==2 {
			result[infoOne[0]] =infoOne[1]
		}
	}
	c.JSON(200, result)
}
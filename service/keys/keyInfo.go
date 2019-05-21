package keys

import (
	"github.com/RXDA/RedisAPI/config"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
)

func Value(c *gin.Context) {
	//获取参数
	key := c.Query("key")  //key name
	stringType := c.Query("type") //string key 类型
	var strType int
	if stringType!="" {
		i, err := strconv.Atoi(stringType)
		if err != nil {
			log.Fatal(err)
		}else{
			strType =i
		}
	}

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
	switch keyType {
	//string ,bitmap,hyperLogLog
	case "string":
		res := getString(key, strType)
		c.JSON(200, res)
	case "hash":
		res := getHashKeys(key)
		c.JSON(200, res)
	case "set":
		res := getSet(key)
		c.JSON(200,res)
	//geo,zset
	case "zset":
	case "list":
	case "none":
	}

}

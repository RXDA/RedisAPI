package hash

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"github.com/RXDA/RedisAPI/config"
	"strconv"
)

var rdc redis.Conn


//获取hash的key列表
func Hkeys(c *gin.Context) {
	key := c.Query("key")
	rdc = config.Conn.RedisPool.Get()
	res, err := redis.Strings(rdc.Do("HKEYS", key))
	if err != nil {
		c.JSON(500,err)
	}
	c.JSON(200,res)
}



//获取hash的所有key value
func HGetAll(c *gin.Context) {
	key := c.Query("key")
	rdc = config.Conn.RedisPool.Get()
	res, err := redis.StringMap(rdc.Do("HGETALL", key))
	if err != nil {
		c.JSON(500,err)
	}
	c.JSON(200,res)
}

//获取hash的某个value
func HGet(c *gin.Context){
	key := c.Query("key")
	field := c.Query("field")
	rdc = config.Conn.RedisPool.Get()
	res, err := redis.String(rdc.Do("HGET", key, field))
	if err != nil {
		c.JSON(500,err)
	}
	c.JSON(200,res)
}

//计数器加整数
func Hincrby(c *gin.Context) {
	key := c.Query("key")
	field := c.Query("field")
	step := c.Query("step")
	stepNum, err:=strconv.Atoi(step)
	if err !=nil{
		c.JSON(500,err)
	}
	rdc = config.Conn.RedisPool.Get()
	res, err := redis.Int64(rdc.Do("HINCRBY", key, field, stepNum))
	if err != nil {
		c.JSON(500,err)
	}
	c.JSON(200,res)
}
//计数器加浮点数
func HincrbyFloat(c *gin.Context) {
	key := c.Query("key")
	field := c.Query("field")
	step := c.Query("step")
	stepNum, err:=strconv.ParseFloat(step, 64)

	rdc = config.Conn.RedisPool.Get()
	res, err := redis.Float64(rdc.Do("HINCRBYFLOAT", key, field, stepNum))
	if err != nil {
		c.JSON(500,err)
	}
	c.JSON(200,res)
}

//hash长度
func Hlen(c *gin.Context){
	key := c.Query("key")
	rdc = config.Conn.RedisPool.Get()
	res, err := redis.Int(rdc.Do("HGETALL", key))
	if err != nil {
		c.JSON(500,err)
	}
	c.JSON(200,res)
}


//hash长度
func HMGet(c *gin.Context){
	key := c.Query("key")
	rdc = config.Conn.RedisPool.Get()
	res, err := redis.Int(rdc.Do("HGETALL", key))
	if err != nil {
		c.JSON(500,err)
	}
	c.JSON(200,res)
}

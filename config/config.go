package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Connection struct {
	DBPool *gorm.DB
	RedisPool *redis.Pool
}


func newRedisPool(addr string) *redis.Pool {
	if _,err:=redis.Dial("tcp", addr);err!=nil{
		log.Fatal(err)
	}
	return &redis.Pool{
		MaxIdle: 30,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}


var Conn Connection


func init(){
	func (env string) {
		config := viper.New()
		config.SetConfigType("yaml")
		config.SetConfigName(env)
		config.AddConfigPath("./config/")
		config.AddConfigPath("config")

		err := config.ReadInConfig()
		if err != nil {
			log.Fatal(err.Error())
		}
		//redis pool
		Conn.RedisPool = newRedisPool(config.GetString("redis.url"))
	}("develop")
}

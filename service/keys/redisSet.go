package keys

import (
	"github.com/RXDA/RedisAPI/config"
	"github.com/gomodule/redigo/redis"
	"log"
)

//get set members
func getSet(key string) []string{
	conn := config.Conn.RedisPool.Get()
	res, err := redis.Strings(conn.Do("SMEMBERS", key))
	if err != nil {
		log.Fatal(err)
	}
	return res
}
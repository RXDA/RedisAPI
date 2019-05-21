package keys

import (
	"github.com/RXDA/RedisAPI/config"
	"github.com/gomodule/redigo/redis"
	"log"
)

const (
	STRING = iota
	BITMAP
	HYPERLOGLOG
)

func getString(key string, keyType int) interface{} {
	conn := config.Conn.RedisPool.Get()
	switch keyType {
	case STRING:
		res, err := redis.String(conn.Do("GET", key))
		if err != nil {
			log.Fatal(err)
		}
		return res
	case BITMAP:
		bitmap, err := redis.Bytes(conn.Do("GET", key))
		if err != nil {
			log.Fatal(err)
		}
		str :=getBitSet(bitmap)
		return str
	case HYPERLOGLOG:
		pfCount, err := redis.Int(conn.Do("PFCOUNT", key))
		if err != nil {
			log.Fatal(err)
		}
		return pfCount
	default:
		return nil
	}
}


func hasBit(n byte, pos uint) bool {
	val := n & (1 << pos)
	return val > 0
}


func getBitSet(redisResponse []byte) string {
	bitset := make([]bool, len(redisResponse)*8)

	for i := range redisResponse {
		for j:=7; j>=0; j-- {
			bitN := uint(i*8+(7-j))
			bitset[bitN] = hasBit(redisResponse[i], uint(j))
		}
	}
	result := ""
	for _,v :=range bitset  {
		if v {
			result +="1"
		}else{
			result +="0"
		}
	}

	return result
}
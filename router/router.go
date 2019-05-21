package router

import "github.com/gin-gonic/gin"
import  "github.com/RXDA/RedisAPI/service/hash"

func Router() *gin.Engine {
	r := gin.Default()

	hashGroup := r.Group("/hash")
	{
		hashGroup.GET("hkeys", hash.Hkeys)
		hashGroup.GET("hgetall", hash.HGetAll)
		hashGroup.GET("hget", hash.HGet)
		hashGroup.GET("hincrby", hash.Hincrby)
		hashGroup.GET("hincrbyfloat", hash.HincrbyFloat)
		hashGroup.GET("hlen", hash.Hlen)
		hashGroup.GET("hkeys", hash.Hkeys)
		hashGroup.GET("hkeys", hash.Hkeys)

	}

	return r
}

package main

import (
	"coindesk/cache"
	"github.com/gin-gonic/gin"
)

func main() {

	var (
		redisClient, err = cache.NewRedisClient()
		cs               = NewCryptoPriceService(redisClient)
	)

	if err != nil {
		logger.Error(err.Error())
	}

	router := gin.Default()
	routerGroup := router.Group("/coindesk")
	CryptoHttpTransport(routerGroup, cs)
	err = router.Run("localhost:8080")
	if err != nil {
		logger.Error(err.Error())
	}

}

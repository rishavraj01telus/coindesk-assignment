package main

import (
	"coindesk/cache"
	"coindesk/service"
	"coindesk/transport"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

var logger, _ = zap.NewProduction()

func main() {

	var (
		redisClient, err = cache.NewRedisClient()
		cacheStorage     = cache.NewCacheStorage(redisClient)
		cs               = service.NewCryptoPriceService(&cacheStorage)
	)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	router := gin.Default()
	routerGroup := router.Group("/coindesk")
	transport.CryptoHttpTransport(routerGroup, cs)
	err = router.Run("localhost:8080")
	if err != nil {
		logger.Error(err.Error())
	}

}

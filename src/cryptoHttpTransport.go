package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CryptoHttpTransport(routerGroup *gin.RouterGroup, cs *CryptoService) {
	routerGroup.GET("/price", cryptoPriceHandler(cs))
}

var cs CryptoService

func cryptoPriceHandler(cs *CryptoService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		price, err := cs.CryptoPriceService()

		if err != nil {
			logger.Error(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": price,
		})
	}
}

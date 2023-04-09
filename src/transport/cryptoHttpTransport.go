package transport

import (
	s "coindesk/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var logger, _ = zap.NewProduction()

func CryptoHttpTransport(routerGroup *gin.RouterGroup, cs *s.CryptoService) {
	routerGroup.GET("/price", cryptoPriceHandler(cs))
}

func cryptoPriceHandler(cs *s.CryptoService) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		price, err := cs.CryptoPriceService(ctx)

		if err != nil {
			logger.Error(err.Error())
			ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": price,
		})
	}
}

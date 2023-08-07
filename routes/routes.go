package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"real_my_tiktok/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "test ok",
		})
	})
	return r
}

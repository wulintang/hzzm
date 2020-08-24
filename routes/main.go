package routes

import (
	"github.com/gin-gonic/gin"
	"hzzm/config"
	"hzzm/controller"
	"hzzm/middlewares"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 设置gin模式
	gin.SetMode(config.GinMode)

	r := gin.Default()
	r.Use(middlewares.Cors())

	r.GET("/ping", func(c *gin.Context) { //服务健康检查
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})


	v1 := r.Group("/v1")
	{
		v1.GET("/list",controller.Paginate)
	}

	return r
}
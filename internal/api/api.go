package api

import (
	"L0/internal/caching"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
	Router *gin.Engine
	cache  *caching.Cache
}

func (a *Api) InitRouter(cache *caching.Cache) {
	a.cache = cache
	a.Router = gin.Default()

	a.Router.Use(CORSMiddleware())

	a.Router.GET("/order/", a.getOrder)
}

func (a *Api) getOrder(c *gin.Context) {
	id := c.Query("id")
	order, err := a.cache.Get(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, order)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		c.Next()
	}
}

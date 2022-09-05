package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ConfigureRouter() *gin.Engine {
	router := gin.New()

	test := router.Group("/test")
	{
		test.GET("/", h.test)
		test.GET("/say-hello", h.sayHello)
	}

	ws := router.Group("/ws")
	{
		ws.GET("/", h.connect)
		ws.GET("/echo", h.echo)
	}

	return router
}

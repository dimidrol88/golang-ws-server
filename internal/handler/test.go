package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) sayHello(c *gin.Context) {
	logrus.Infof("sayHello from:" + c.Request.RemoteAddr)
}

func (h *Handler) test(c *gin.Context) {
	logrus.Infof("test from:" + c.Request.RemoteAddr)
}

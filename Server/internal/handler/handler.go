package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	// there will be services soon...
}

func NewHandler() *Handler {
	// there will be services soon...
	return &Handler{}
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"status" : "ok"})
}
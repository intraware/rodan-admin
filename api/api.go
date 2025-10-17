package api

import (
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/intraware/rodan-admin/internal/utils/values"
)

func LoadRoutes(r *gin.Engine) {
	apiRouter := r.Group("/api")

	shared.Init(values.GetConfig())
	apiRouter.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "pong"})
	})
}
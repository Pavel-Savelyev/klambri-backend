package routes

import (
	"klambri-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/playbook", func(c *gin.Context) {
		handlers.PlaybookHandler(c)
	})
}

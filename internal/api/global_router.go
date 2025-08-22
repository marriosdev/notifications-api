package api

import (
	"github.com/gin-gonic/gin"
	"github.com/marriosdev/export-api/internal/domain/notification"
)

func RegisterRouter(r *gin.Engine) {
	notification.RegisterRouter(r)
}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/marriosdev/export-api/internal/database"
	"github.com/marriosdev/export-api/internal/domain/notification"
)

func RegisterRouter(r *gin.Engine) {
	_, mongodatabase, _ := database.NewMongoDB()
	notificationRepository := notification.NewRepository(mongodatabase)
	notificationService := notification.NewService(notificationRepository)
	setupNotificationRoutes(r, notification.NewHandler(notificationService))
}

func setupNotificationRoutes(r *gin.Engine, handler *notification.Handler) {
	r.GET("/notifications", handler.GetAllNotifications)
}

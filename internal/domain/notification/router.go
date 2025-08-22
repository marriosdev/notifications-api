package notification

import (
	"github.com/gin-gonic/gin"
	"github.com/marriosdev/export-api/internal/database"
)

func RegisterRouter(r *gin.Engine) {
	_, mongodatabase, _ := database.NewMongoDB()
	notificationRepository := NewRepository(mongodatabase)
	notificationService := NewService(notificationRepository)
	handler := NewHandler(notificationService)

	r.GET("/notifications", handler.GetAllNotifications)
	r.GET("/notifications/:id", handler.GetById)
	r.POST("notification", handler.SaveNotification)
}

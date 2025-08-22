package notification

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllNotifications(c *gin.Context) {
	ctx := c.Request.Context()
	notifications, err := h.service.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notifications)
}

func (h *Handler) GetById(c *gin.Context) {
	notification, err := h.service.repo.GetById(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notification)
}

func (h *Handler) SaveNotification(c *gin.Context) {
	var notification Notification

	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	n, err := h.service.SaveNotification(c, notification)

	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, n)
}

package message

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	service *MessageService

	logger *slog.Logger
}

func NewHandler(service *MessageService, logger *slog.Logger) *MessageHandler {
	return &MessageHandler{
		service: service,
		logger:  logger,
	}
}

type Message struct {
	Content string `json:"content"`
}

func (h MessageHandler) Receiver(c *gin.Context) {
	var msg Message
	if err := c.BindJSON(&msg); err != nil {
		h.logger.Debug("failed bind json: " + err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.SaveMsg(msg.Content); err != nil {
		h.logger.Error("failed save msg: " + err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

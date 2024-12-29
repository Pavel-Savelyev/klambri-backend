package handlers

import (
	natsclient "klambri-backend/internal/nats-client"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PlaybookHandler(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body: " + err.Error()})
		return
	}

	if err := natsclient.Publish("playbookQueue", jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send to NATS: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "PlaybookConfig sent successfully"})
}

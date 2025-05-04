package pkg

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

type auditLog struct {
	Action     string
	IpAddress  string
	UserAgent  string
	StatusCode int
	CreatedAt  time.Time
}

func LogAction(c *gin.Context, action string, statusCode int) {
	log := auditLog{
		Action:     action,
		IpAddress:  c.ClientIP(),
		UserAgent:  c.Request.UserAgent(),
		StatusCode: statusCode,
		CreatedAt:  time.Now(),
	}

	logJSON, _ := json.Marshal(log)
	logFile, _ := os.OpenFile("audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer logFile.Close()
	logFile.Write(logJSON)
	logFile.WriteString("\n")
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"twiteer/config"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := config.GetSession(c.Request)
		if err != nil || session.Values["userID"] == nil {
			log.Println("[AUTH] Unauthorized access attempt")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Optionally store userID in context for later use
		c.Set("userID", session.Values["userID"])
		c.Next()
	}
}

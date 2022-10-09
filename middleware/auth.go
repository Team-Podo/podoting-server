package middleware

import (
	"context"
	"fmt"
	"github.com/Team-Podo/podoting-server/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if idToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID token not available"})
		c.Abort()
		return
	}

	//verify token
	token, err := auth.Firebase.VerifyToken(idToken, context.Background())
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}

	c.Set("UUID", token.UID)
	c.Next()
}

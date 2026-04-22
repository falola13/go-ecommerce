package middleware

import (
	"net/http"
	"strings"

	token "github.com/falola13/go-ecommerce/tokens"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.GetHeader("Authorization")
		if ClientToken != "" {
			const bearerPrefix = "Bearer "
			if strings.HasPrefix(ClientToken, bearerPrefix) {
				ClientToken = strings.TrimSpace(strings.TrimPrefix(ClientToken, bearerPrefix))
			}
		}
		if ClientToken == "" {
			ClientToken = c.GetHeader("token")
		}
		if ClientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(ClientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_Name)
		c.Set("last_name", claims.Last_Name)
		c.Set("phone", claims.Phone)
		c.Set("uid", claims.Uid)
		c.Next()
	}

}

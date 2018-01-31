package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(app *App) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Println(authHeader)
		if authHeader == "" {
			app.responseHandler.jwtAbort(c, "Authorization头不存在")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			app.responseHandler.jwtAbort(c, "Authorization头无效")
			return
		}

		claims, err := ParseToken(parts[1])
		if err != nil {
			app.responseHandler.jwtAbort(c, "无效的Token")
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			app.responseHandler.jwtAbort(c, "Token已过期")
			return
		}

		user := User{}
		app.db.First(&user, claims.UserID)

		if user.ID != claims.UserID {
			app.responseHandler.jwtAbort(c, "无效的Token")
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

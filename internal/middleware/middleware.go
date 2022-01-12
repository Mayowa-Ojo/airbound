package middleware

import (
	"airbound/internal"
	"airbound/internal/auth"
	"airbound/internal/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeaderPrefix = "Bearer "
)

type AuthorizationHeader struct {
	Authorization string `header:"Authorization"`
}

func AuthenticateUser(cfg *config.Config) gin.HandlerFunc {
	// check auth status
	return func(c *gin.Context) {
		var h AuthorizationHeader

		if err := c.ShouldBindHeader(&h); err != nil {
			c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{
				"status":  "error",
				"message": fmt.Sprint(err),
				"data":    nil,
			})
			return
		}

		t := strings.Split(h.Authorization, AuthorizationHeaderPrefix)
		if len(t) != 2 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "missing/malformed token",
				"data":    nil,
			})
			return
		}

		token := t[1]

		tokenData, err := auth.VerifyToken(token, cfg)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": fmt.Sprint(err),
				"data":    nil,
			})
			return
		}

		tokenData.Token = token
		c.Set("token-data", tokenData)

		c.Next()
	}
}

func AuthorizeUser(permission string) gin.HandlerFunc {
	// check permissions
	return func(c *gin.Context) {
		td, ok := c.Get("token-data")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "missing/malformed token",
				"data":    nil,
			})
			return
		}

		tokenData := td.(*auth.TokenData)

		if tokenData.User.Role.Name != tokenData.Role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  "error",
				"message": "action not allowed",
				"data":    nil,
			})
			return
		}

		var permissions []string
		for _, v := range tokenData.User.Role.Permissions {
			permissions = append(permissions, v.Permission)
		}

		hasPermission := internal.ContainsString(permission, permissions)
		if !hasPermission {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  "error",
				"message": "action not allowed",
				"data":    nil,
			})
			return
		}

		c.Next()
	}
}

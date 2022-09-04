package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labscool/mb-appointment-system/internal/feature/users"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header found"})
			return
		}
		tokenString := authHeader[len(BearerSchema):]

		if token, err := users.ValidateToken(tokenString); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			if claims, ok := token.Claims.(*users.JWTClaim); !ok {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "The token has invalid claims"})
			} else {
				if token.Valid {
					ctx.Set("username", claims.Username)
					ctx.Set("id", claims.ID)
				} else {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				}

			}
		}
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"hunch-api/src/database/model"
	"hunch-api/src/util/token"
	"net/http"
)

func JwtAuthMiddleWare(validRoles []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := token.ValidateToken(context)

		if err != nil {
			context.String(http.StatusUnauthorized, "unauthorized")
			context.Abort()
			return
		}

		tokenType, err := token.ExtractTokenType(context)

		if err != nil || tokenType != "access" {
			context.String(http.StatusUnauthorized, "unauthorized")
			context.Abort()
			return
		}

		id, err := token.ExtractTokenID(context)

		if err != nil {
			context.String(http.StatusUnauthorized, "unauthorized")
			context.Abort()
			return
		}

		user, err := model.GetUserById(id)

		if err != nil {
			context.String(http.StatusUnauthorized, "unauthorized")
			context.Abort()
			return
		}

		if len(validRoles) > 0 {
			valid := false

			for _, validRole := range validRoles {
				for _, userRole := range user.Roles {
					if userRole.Name != validRole {
						continue
					}
					valid = true
				}
			}

			if !valid {
				context.String(http.StatusUnauthorized, "unauthorized")
				context.Abort()
				return
			}
		}

		context.Next()
	}
}

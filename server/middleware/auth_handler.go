package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"server/db/mongo"
	"server/db/mongo/mongodb"
	"server/utils"
)

func AuthHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := context.Cookie("jwt")
		if err != nil {
			utils.RespondError(context, 400, errors.New("no authentication found"))
			context.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			utils.RespondError(context, 400, err)
			context.Abort()
			return
		}
		user, err := mongo.Queries().GetUserByTokenAndId(context.Request.Context(), mongodb.GetUserByTokenAndIdParams{
			ID:          claims.ID,
			AccessToken: token,
		})
		if err != nil {
			utils.RespondError(context, 400, err)
			context.Abort()
			return
		}
		context.Set("user", user)
		context.Next()
	}
}

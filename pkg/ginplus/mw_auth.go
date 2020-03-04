package ginplus

import (
	"context"
	"doudou/pkg/auth"
	"doudou/pkg/errors"

	"github.com/gin-gonic/gin"
)

var AuthS = auth.DefaultAuthServer()

// AuthMiddleware 用户授权中间件
func AuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID string
		if t := GetToken(c); t != "" {
			uid, err := AuthS.VertifyToken(context.TODO(), t)
			if err != nil {
				if err.Error() == errors.ErrInvalidToken.Error() {
					ResError(c, errors.ErrInvalidToken)
					return
				}
				ResError(c, err)
				return
			}
			userID = uid
		}

		if userID != "" {
			c.Set(UserIDKey, userID)
		}

		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		if userID == "" {
			ResError(c, errors.ErrNoPerm)
		}
	}
}

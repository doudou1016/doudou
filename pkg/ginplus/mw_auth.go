package ginplus

import (
	"doudou/pkg/errors"
	"doudou/pkg/jwtplus"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 用户授权中间件
func AuthMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID string
		if t := GetToken(c); t != "" {
			claims, err := jwtplus.ParseToken(t)
			if err != nil {
				if err.Error() == errors.ErrInvalidToken.Error() {
					ResError(c, errors.ErrInvalidToken)
					return
				}
				ResError(c, err)
				return
			}
			userID = claims.UserId
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

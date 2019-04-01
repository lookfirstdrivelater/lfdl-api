package lfdlapi

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func authMiddleware(db *gorm.DB) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &user{
				UserName: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userName := loginVals.Username
			password := loginVals.Password

			var localUser user
			db.Where("user_name = ?", userName).First(&localUser)

			if userName == localUser.UserName && checkPasswordHash(password, localUser.Password) {
				return &user{
					UserName:  localUser.UserName,
					FirstName: localUser.FirstName,
					LastName:  localUser.LastName,
				}, nil

			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*user); ok {
				var localUser user
				db.Where("user_name = ?", v.UserName).First(&localUser)
				if localUser.AuthGeneral {
					return true
				}
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware, err
}


// this middleware allows us to auth users for only the routes they are authorized for
func groupAuthorizator(group string, authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {

	return func(c *gin.Context) {
		switch group {
		case "general":
			usr, _ := c.Get("id")

			var localUser user
			db.Where("user_name = ?", usr.(*user).UserName).First(&localUser)
			if !localUser.AuthGeneral {
				authMiddleware.Unauthorized(c, http.StatusUnauthorized, authMiddleware.HTTPStatusMessageFunc(jwt.ErrForbidden, c))
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Next()
			return
		case "admin":
			usr, _ := c.Get("id")

			var localUser user
			db.Where("user_name = ?", usr.(*user).UserName).First(&localUser)
			if !localUser.AuthAdmin {
				authMiddleware.Unauthorized(c, http.StatusUnauthorized, authMiddleware.HTTPStatusMessageFunc(jwt.ErrForbidden, c))
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			c.Next()
			return
		default:
			fmt.Println("here")
			//authMiddleware.Unauthorized(c, http.StatusForbidden, authMiddleware.HTTPStatusMessageFunc(jwt.ErrForbidden, c))
		}

	}

}
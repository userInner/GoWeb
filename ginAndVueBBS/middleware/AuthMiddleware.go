package middleware

import (
	"fmt"
	"ginAndVueBBS/common"
	"ginAndVueBBS/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		fmt.Println(claims.UserId)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}

		//用户
		userId := claims.UserId
		user := dao.GetUserIdUser(userId)
		if user.UserId == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

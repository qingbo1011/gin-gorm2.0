package route

import (
	"gin-gorm2.0/api"
	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	// gin.Default()和gin.New()的区别：gin.Default()默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/register", api.UserRegister)
		user.POST("/login", api.Login)
	}
	return r
}

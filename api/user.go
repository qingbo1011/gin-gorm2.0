package api

import (
	"net/http"

	"gin-gorm2.0/request"
	"gin-gorm2.0/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func UserRegister(c *gin.Context) {
	var userRegister request.UserRep
	err := c.ShouldBind(&userRegister)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "json数据解析失败！",
			"error":  err.Error(),
		})
		logging.Info(err)
		return
	}
	res, err := service.UserRegister(userRegister)
	if err != nil {
		c.JSON(res.Status, res)
	}
	c.JSON(http.StatusCreated, res)
}

func Login(c *gin.Context) {

}

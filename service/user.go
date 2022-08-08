package service

import (
	"net/http"

	"gin-gorm2.0/db/mysql"
	"gin-gorm2.0/model"
	"gin-gorm2.0/request"
	"gin-gorm2.0/response"
	"gorm.io/gorm"
)

func UserRegister(register request.UserRep) (response.Response, error) {
	var user model.User
	err := mysql.MysqlDB.Where(&model.User{UserName: register.UserName}).First(&user).Error
	if err == gorm.ErrRecordNotFound { // ErrRecordNotFound错说明数据库没有记录，即UserName没有重复，可以注册
		user.UserName = register.UserName
		err := user.SetPassword(register.Password)
		if err != nil {
			return response.Response{
				Status: http.StatusInternalServerError,
				Msg:    "密码加密出错！",
				Error:  err.Error(),
			}, err
		}
		err = mysql.MysqlDB.Create(&user).Error
		if err != nil {
			return response.Response{
				Status: http.StatusInternalServerError,
				Msg:    "数据库添加数据出错！",
				Error:  err.Error(),
			}, err
		}
		return response.Response{
			Status: http.StatusCreated,
			Msg:    "用户注册成功！",
		}, err
	} else if err != nil {
		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "数据库内部出错！",
			Error:  err.Error(),
		}, err
	}
	return response.Response{
		Status: http.StatusBadRequest,
		Msg:    "用户名重复，注册失败！",
	}, err
}

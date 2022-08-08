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
		Status: http.StatusForbidden,
		Msg:    "用户名重复，注册失败！",
	}, err
}

func UserLogin(login request.UserRep) (response.Response, error) {
	var user model.User
	err := mysql.MysqlDB.Where(&model.User{UserName: login.UserName}).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.Response{
				Status: http.StatusBadRequest,
				Msg:    "该用户不存在，请先注册！",
				Error:  err.Error(),
			}, err
		}
		// 不是用户不存在却还是出错，其他不可抗拒的因素
		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "查询数据库出现错误！",
			Error:  err.Error(),
		}, err
	}
	// 用户从数据库中找到了，检验密码
	ok, err := user.CheckPassword(login.Password)
	if err != nil {
		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "登录失败！",
			Error:  err.Error(),
		}, err
	}
	if !ok {
		return response.Response{
			Status: http.StatusForbidden,
			Msg:    "密码错误，登录失败！",
		}, err
	}
	// 登录成功，这里就省略签发token了
	return response.Response{
		Status: http.StatusOK,
		Msg:    "登录成功！",
		Data:   map[string]string{"token": "这里就模拟一个token"},
	}, nil
}

package UserService

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"modulename/basic_homework/homework_for_redrock/test/业务/model"
	"modulename/basic_homework/homework_for_redrock/test/业务/service"
	"modulename/basic_homework/homework_for_redrock/test/业务/util"
)

func Register(c *gin.Context) {
	userName := c.PostForm("name")
	password := c.PostForm("password")
	if userName == "" || password == "" {
		util.RespParamErr(c)
		return
	}

	u, err := service.SearchUserByUserName(userName)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RespInternalErr(c)
		return
	}

	if u.UserName != "" {
		util.NomalErr(c, 300, "账户已存在")
		return
	}

	err = service.CreateUser(model.User{
		UserName: userName,
		Password: password,
	})
	if err != nil {
		util.RespInternalErr(c)
		return
	}
	util.RespOK(c)
}

func Login(c *gin.Context) {
	userName := c.PostForm("name")
	password := c.PostForm("password")
	if userName == "" || password == "" {
		util.RespParamErr(c)
		return
	}

	u, err := service.SearchUserByUserName(userName)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NomalErr(c, 300, "用户不存在")
		} else {
			log.Printf("search user error:%v", err)
			util.RespInternalErr(c)
			return
		}
		return
	}
	if u.Password != password {
		util.NomalErr(c, 200, "密码错误")
		return
	}
	c.SetCookie("name", userName, 0, "", "/", false, false)
}

package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) GetUser() {
	fmt.Println(beego.AppConfig.String("dev::httpport"))
	u.Ctx.WriteString("test")
}

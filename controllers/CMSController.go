package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type CMSController struct {
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (this *CMSController) StaticBlock() {
	//fmt.Println(beego.AppConfig.String("dev::httpport"))
	fmt.Println(this.GetString(":key"))
	fmt.Println(this.Input().Get("key"))
	this.Ctx.WriteString("StaticBlock")
}

// @router /all/:key [get]
func (this *CMSController) AllBlock() {
	//fmt.Println(beego.AppConfig.String("dev::httpport"))
	fmt.Println(this.GetString(":key"))
	fmt.Println(strconv.Atoi(this.Input().Get("key")))
	this.Ctx.WriteString("AllBlock")
}

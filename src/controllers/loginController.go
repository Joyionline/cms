package controllers

import (
	"cms/src/common"
	"cms/src/service"
	"strconv"

	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

/**
进入登录页面
*/
func (this *LoginController) Tologin() {
	//展示页面
	this.show("common/loginPage.html")
}

/**
登陆
*/
func (this *LoginController) Login() {
	//接收页面传递过来的参数(账号和密码)
	accout := this.GetString("accout")
	password := this.GetString("password")
	//调用加密方法进行MD5加密
	encodePwd := common.EncodeMessageMd5(password)
	//调用业务层中登录鉴权方法，对当前登录用户的权限进行校验
	if admusers, err := service.AdmUserService.Authentication(accout, encodePwd); err != nil {
		this.jsonResult(err.Error())
	} else {
		token := strconv.FormatInt(admusers.Id, 10) + "|" + accout + "|" + this.getClientIp()
		token = common.EncryptAes(token)
		this.Ctx.SetCookie("token", token, 0)
		this.jsonResult(SUCCESS)
	}
}

/**
退出登陆
*/
func (this *LoginController) Loginout() {
	this.Ctx.SetCookie("token", "", 0)
	this.redirect(beego.URLFor("LoginController.Tologin"))
}

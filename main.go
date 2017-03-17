package main

import (
	_ "cms/src/routers"
	_ "cms/src/service"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogger("console", "")
	//执行beego框架运行环境
	beego.Run()
}

package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lxxccc/tfxingApiProxy/models"
	_ "github.com/lxxccc/tfxingApiProxy/routers"
)

func main() {
	beego.Run()
}

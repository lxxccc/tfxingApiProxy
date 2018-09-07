package main

import (
	"github.com/astaxie/beego"
	_ "github.com/tfxingApiProxy/models"
	_ "github.com/tfxingApiProxy/routers"
)

func main() {
	beego.Run()
}

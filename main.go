package main

import (
	"huodong/models"
	_ "huodong/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetLevel(beego.LevelInformational)
	go models.Run()

	beego.Run()

}

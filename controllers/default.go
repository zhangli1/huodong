package controllers

import (
	"fmt"
	"huodong/models"
	glib "lib"

	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type RequestCount struct {
	Sign      string
	Num       int
	Timestamp int
}

var SameRequestLimit = make(map[string]RequestCount, 0)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Add() interface{} {
	r := c.Ctx.Request
	fmt.Println(c.Ctx.Input.IP(), r.Header["User-Agent"])

	Id := c.Input().Get("id")
	fmt.Println("add: " + Id)

	//先清除过期的数据
	if len(SameRequestLimit) > 0 {
		for i, v := range SameRequestLimit {
			if v.Timestamp < glib.DateToTimestamp("Y-m-d", fmt.Sprintf("%s-%s-%s", time.Now().Year(), time.Now().Month(), time.Now().Day())) {
				delete(SameRequestLimit, i)
			}
		}
	}

	var md5Sign string
	md5Sign = models.Md5(fmt.Sprintf("%s,%d", r.Header["User-Agent"], Id))

	if _, ok := SameRequestLimit[md5Sign]; ok && SameRequestLimit[md5Sign].Num >= 5 {
		fmt.Println("超出限制", SameRequestLimit[md5Sign])
		return nil
	}
	request := RequestCount{}
	if _, ok := SameRequestLimit[md5Sign]; ok {
		request.Sign = md5Sign
		request.Num = SameRequestLimit[md5Sign].Num + 1
	} else {
		request.Sign = md5Sign
		request.Num = 1
	}
	request.Timestamp = glib.GetCurrentTime()
	SameRequestLimit[md5Sign] = request

	fmt.Println(SameRequestLimit)

	IntId, _ := strconv.Atoi(Id)
	models.AddCount(IntId)
	//c.TplName = "index2.tpl"

	return nil
}

func (c *MainController) Search() {
	var QiNiuPath string
	QiNiuPath = beego.AppConfig.String("imgPath")
	//QiNiuPath = "http://or84xoiz8.bkt.clouddn.com/"
	r := c.Ctx.Request
	fmt.Println(c.Ctx.Input.IP(), r.Header["User-Agent"])
	var mapData map[string]interface{}
	mapData = make(map[string]interface{}, 4)
	data := models.SearchVote()

	maxData := models.GetMaxCount()

	mapData["Id"] = maxData[0].Id
	mapData["Name"] = maxData[0].Name
	mapData["Img"] = maxData[0].Img
	mapData["Count"] = maxData[0].Count

	c.Data["MaxVote"] = mapData
	c.Data["Votes"] = data
	c.Data["QiNiuPath"] = "http://" + QiNiuPath
	c.TplName = "index.tpl"
}

//获取聊天消息，默认根据传上来的id取后面的数据
func (c *MainController) GetLastMessage() {
	Id := c.Input().Get("lastId")
	//fmt.Println(Id)
	retData := models.SearchLastMessage(Id)

	c.Data["json"] = retData
	c.ServeJSON()
}

func (c *MainController) AddMessage() {
	r := c.Ctx.Request
	fmt.Println("AddMessage", c.Ctx.Input.IP(), r.Header["User-Agent"])
	Content := c.Input().Get("content")
	models.AddMessage(Content, c.Ctx.Input.IP())
}

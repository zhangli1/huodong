package models

import (
	"fmt"
	glib "lib"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Vote struct {
	Id    int
	Name  string
	Count int
	Img   string
}

type Message struct {
	Id         int
	Content    string
	Ip         string
	CreateTime string
}

var o orm.Ormer
var GlobalMessageListData []Message

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysql"))
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Vote))
	orm.RegisterModel(new(Message))
	o = InitDb("default")
}

func InitDb(DbName string) orm.Ormer {
	o := orm.NewOrm()
	o.Using(DbName) // 默认使用 default，你可以指定为其他数据库
	return o
}

func AddCount(Id int) {
	vote := Vote{Id: Id}
	_ = o.Read(&vote)
	addCount := vote.Count
	vote.Count = addCount + 1

	o.Update(&vote, "Count")
}

func SearchVote() []Vote {
	var votes []Vote
	num, err := o.Raw("select * from vote").QueryRows(&votes)
	if err == nil && num > 0 {
		return votes
	}
	return votes
}

func GetMaxCount() []Vote {
	var votes []Vote
	num, err := o.Raw("select * from vote order by count desc limit 1").QueryRows(&votes)
	if err == nil && num > 0 {
		return votes
	}
	return votes
}

//获取消息
func SearchLastMessage(LastId string) []Message {
	var messages []Message
	if len(GlobalMessageListData) > 0 {
		for _, v := range GlobalMessageListData {
			if v.Id > glib.StringToInt(LastId) {
				messages = append(messages, v)
			}
		}
		return messages
	}

	num, err := o.Raw(fmt.Sprintf("select * from message where id > %s limit 1000", LastId)).QueryRows(&messages)
	if err == nil && num > 0 {
		if len(GlobalMessageListData) < 1 {
			GlobalMessageListData = append(GlobalMessageListData, messages...)
		}
		return messages
	}
	return messages
}

//添加消息
func AddMessage(content string, ip string) {
	createTime := glib.TimestampToDate("", glib.GetCurrentTime())
	var message Message
	message.Content = content
	message.Ip = ip
	message.CreateTime = createTime

	id, err := o.Insert(&message)
	if err == nil {
		fmt.Println(id)
	} else {
		fmt.Println(err)
	}
	message.Id = int(id)
	GlobalMessageListData = append(GlobalMessageListData, message)
}

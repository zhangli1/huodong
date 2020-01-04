package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

func Run() {
	http.Handle("/getMessage", websocket.Handler(getMessage))

	if err := http.ListenAndServe(":10001", nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getMessage(ws *websocket.Conn) {
	//主动推送数据
	/*go func() {
		for {
			//websocket.Message.Send(ws, "主动推送")
			time.Sleep(time.Second * time.Duration(1))
		}
	}()*/

	//有问有答数据
	var err error

	var n = 0
	for {
		var reply string

		//错误10次断开连接
		if n > 10 {
			ws.Close()
			break
		}

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			//fmt.Println(err)
			n++
			continue
		}

		if reply == "" {
			reply = "0"
		}
		json, _ := json.Marshal(SearchLastMessage(reply))

		if err = websocket.Message.Send(ws, string(json)); err != nil {
			//fmt.Println(err)
			continue
		}
		time.Sleep(100 * time.Millisecond)
	}
}

package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/liuchong/chat/src/dbases"
	"github.com/liuchong/chat/src/robot"
	"github.com/liuchong/chat/src/server/logger"
)

//Helper 提示信息
var Helper string = `Commands:
=================================
😉 查询CPU 👉 命令明细：查询+IP+CPU
🚀 查询内存 👉 命令明细：查询+IP+内存
🤖 查询硬盘 👉 命令明细：查询+IP+硬盘
=================================
`

//ServerStart 启动服务
func ServerStart(ddtoken string, appsecret string, papi string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Danger(err)
		}
		if len(data) == 0 {
			logger.Warning("回调参数为空，请检查！")
		}
		var msgObj = new(robot.ReMsg)
		err = json.Unmarshal(data, &msgObj)
		if err != nil {
			logger.Warning("接收Body体转换json失败： %v\n", err)
		}
		if msgObj.Text.Content == "" || msgObj.ChatbotUserID == "" {
			logger.Warning("从钉钉回调过来的内容为空，根据过往的经验，或许重新创建一下机器人，能解决这个问题")
			return
		}

		if len(msgObj.Text.Content) == 1 || msgObj.Text.Content == " 帮助" {
			err = msgObj.Forward(Helper, ddtoken)
			if err != nil {
				logger.Danger(err)
			}
		} else {
			newmsg := ProcessReceive(msgObj.Text.Content, papi)
			err = msgObj.Forward(newmsg, ddtoken)
			if err != nil {
				logger.Danger(err)
			}
			logger.Warning(fmt.Sprintf("dingtalk callback parameters: %#v", msgObj))

		}
	}
	server := &http.Server{
		Addr:    ":8081",
		Handler: http.HandlerFunc(handler),
	}
	logger.Info("Start Listen On", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		logger.Danger(err)
	}
}

//ProcessReceive 创建一个函数
func ProcessReceive(ReceiveMsg string, papi string) string {
	newreceivemsg := &dbases.MonitorMsg{
		Address:     ReceiveMsg[7 : len(ReceiveMsg)-6],
		MonitorItem: ReceiveMsg[len(ReceiveMsg)-6:],
	}

	switch newreceivemsg.MonitorItem {
	case "CPU":
		newreceivemsg.MonitorItem = "cpu_usage_idle"
		return newreceivemsg.QueryResult(papi)
	case "内存":
		newreceivemsg.MonitorItem = "cpu_usage_idle"
		return newreceivemsg.QueryResult(papi)
	case "硬盘":
		newreceivemsg.MonitorItem = "cpu_usage_idle"
		return newreceivemsg.QueryResult(papi)
	default:
		return ""
	}
}

package robot

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//DingTalkMessageForwarder 定义一个接口，其包含了一个 Forward 方法，用于接收消息并将其转发给钉钉机器人。
type DingTalkMessageForwarder interface {
	Forward(message string) error
}

//DingTalkRobotForwarder 定义一个类型，实现了 DingTalkMessageForwarder 接口的 Forward 方法。
type DingTalkRobotForwarder struct {
	accessToken string
}

//Message 定义一个消息结构体
type Message struct {
	MsgType string                 `json:"msgtype"`
	Text    map[string]interface{} `json:"text"`
}

//NewDingTalkRobotForwarder 定义一个函数，创建一个钉钉机器人转发器的实例
func NewDingTalkRobotForwarder(accessToken string) *DingTalkRobotForwarder {
	return &DingTalkRobotForwarder{accessToken: accessToken}
}

//Forward 定义一个方法，我们使用钉钉机器人的 Webhook 地址作为请求 URL，并在 URL 中指定了 access_token 参数，用于身份验证。
func (f *DingTalkRobotForwarder) Forward(message string) error {
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + f.accessToken
	msg := Message{
		MsgType: "text",
		Text: map[string]interface{}{
			"content": message,
		},
	}
	jsonStr, _ := json.Marshal(msg)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	///Content-Type 头部也需要设置为 application/json，以确保消息能够正确被解析。
	req.Header.Set("Content-Type", "application/json")

	//使用 http.Client 发送 POST 请求，将消息发送给钉钉机器人。
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

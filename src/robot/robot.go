package robot

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//ReMsg 接收的消息体模板
type ReMsg struct {
	ConversationID string `json:"conversationId"`
	AtUsers        []struct {
		DingtalkID string `json:"dingtalkId"`
	} `json:"atUsers"`
	ChatbotUserID             string `json:"chatbotUserId"`
	MsgID                     string `json:"msgId"`
	SenderNick                string `json:"senderNick"`
	IsAdmin                   bool   `json:"isAdmin"`
	SessionWebhookExpiredTime int64  `json:"sessionWebhookExpiredTime"`
	CreateAt                  int64  `json:"createAt"`
	ConversationType          string `json:"conversationType"`
	SenderID                  string `json:"senderId"`
	ConversationTitle         string `json:"conversationTitle"`
	IsInAtList                bool   `json:"isInAtList"`
	SessionWebhook            string `json:"sessionWebhook"`
	Text                      Text   `json:"text"`
	RobotCode                 string `json:"robotCode"`
	Msgtype                   string `json:"msgtype"`
}

//Text 消息内容
type Text struct {
	Content string `json:"content"`
}

//Message 定义一个消息结构体
type Message struct {
	MsgType string                 `json:"msgtype"`
	Text    map[string]interface{} `json:"text"`
}

//Forward 定义一个方法，我们使用钉钉机器人的 Webhook 地址作为请求 URL，并在 URL 中指定了 access_token 参数，用于身份验证。
func (r ReMsg) Forward(message string, accessToken string) error {
	url := "https://oapi.dingtalk.com/robot/send?access_token=" + accessToken
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

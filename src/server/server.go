package server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/liuchong/chat/src/robot"
	"github.com/liuchong/chat/src/server/logger"
)

//HeaderInfo 消息头
type HeaderInfo struct {
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
}

//DeployTimestamp 时间戳
func (hi *HeaderInfo) DeployTimestamp(ts string) int64 {
	tsNum, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		panic(err)
	}
	now := time.Now().Unix()
	diff := now - tsNum
	return diff
}

//DeploySign 获得签名
func (hi *HeaderInfo) DeploySign(ts string, secret string) string {
	signNatureString := ts + "\n" + secret
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(signNatureString))
	snData := h.Sum(nil)
	snNature := base64.StdEncoding.EncodeToString(snData)
	return snNature

}

//ServerStart 启动服务
func ServerStart(ddtoken string, appsecret string) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Danger(err)
		}
		if len(data) == 0 {
			logger.Warning("回调参数为空，请检查！")
		}
		var msgObj = new(robot.ReceiveMsg)
		err = json.Unmarshal(data, &msgObj)
		if err != nil {
			logger.Warning("接收Body体转换json失败： %v\n", err)
		}
		if msgObj.Text.Content == "" || msgObj.ChatbotUserID == "" {
			logger.Warning("从钉钉回调过来的内容为空，根据过往的经验，或许重新创建一下机器人，能解决这个问题")
			return
		}

		if len(msgObj.Text.Content) == 1 || msgObj.Text.Content == " 帮助" {

			err = robot.Forward("TEST Hello, World!", ddtoken)
			if err != nil {
				logger.Danger(err)
			}
		} else {
			logger.Info(fmt.Sprintf("dingtalk callback parameters: %#v", msgObj))
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

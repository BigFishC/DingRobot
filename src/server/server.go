package server

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"time"
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

//Start 启动服务
func Start() {
	appSecret := "xxxxxxxxxxxxxxxxxx"
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		headerTimeStr := r.Header.Get("timestamp")
		headerSign := r.Header.Get("sign")
		headerMessage := &HeaderInfo{
			Timestamp: headerTimeStr,
			Sign:      headerSign,
		}
		diffResult := headerMessage.DeployTimestamp(headerMessage.Timestamp)
		snResult := headerMessage.DeploySign(headerMessage.Timestamp, appSecret)
		fmt.Printf("timestamp:%v \n sign: %v", diffResult, snResult)
		if diffResult < 3600 && snResult == headerMessage.Sign {
			w.Write([]byte("Hello World!"))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})

	http.ListenAndServe(":8081", nil)
}

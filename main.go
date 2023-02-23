package main

import (
	// dingrobot "github.com/liuchong/chat/src/robot"
	"github.com/liuchong/chat/src/server"
)

func main() {

	// forwarder := dingrobot.NewDingTalkRobotForwarder("7d59776de9dee9697bc2c8ea0da2dd777e0303b21049e5c6f643523bc606a2f6")
	// err := forwarder.Forward("TEST Hello, World!")
	// if err != nil {
	// 	panic(err)
	// }
	server.Start()
}

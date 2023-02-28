package robot

//ProcessRequest 处理请求
func ProcessRequest(rmsg ReMsg, ddtoken string) error {
	switch rmsg.Text.Content {
	case " 老莫，我想吃鱼了！":
		rmsg.Forward("老莫，我想吃鱼了！", ddtoken)
	}
	return nil
}

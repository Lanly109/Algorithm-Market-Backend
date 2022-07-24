package serializer

// output 输出序列化器
type Output struct {
	Text string `json:"text"`
	Img  string `json:"img"`
}

// BuildOutput 序列化输出
func BuildOutput(text string, img string) Output {
	return Output{
		Text: text,
		Img:  img,
	}
}

// BuildOutputResponse 序列化输出响应
func BuildOutputResponse(text string, img string) Response {
	return Response{
		Data: BuildOutput(text, img),
	}
}

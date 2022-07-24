package serializer

import "singo/model"

// output 输出序列化器
type Output struct {
	InputId uint   `json:"input_id"`
	Type    string `json:"type"`
	Output  string `json:"output"`
}

// BuildOutput 序列化输出
func BuildOutput(input model.Input) Output {
	return Output{
		InputId: input.ID,
	}
}

// BuildOutputResponse 序列化输出响应
func BuildOutputResponse(input model.Input) Response {
	return Response{
		Data: BuildOutput(input),
	}
}

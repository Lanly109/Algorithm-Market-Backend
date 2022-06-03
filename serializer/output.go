package serializer

import "singo/model"

// output 输出序列化器
type Output struct {
	Type   string `json:"type"`
	Output string `json:"output"`
}

// BuildOutput 序列化输出
func BuildOutput(output model.Output) Output {
	return Output{
		Type:   output.Type,
		Output: output.Data,
	}
}

// BuildOutputResponse 序列化输出响应
func BuildOutputResponse(output model.Output) Response {
	return Response{
		Data: BuildOutput(output),
	}
}

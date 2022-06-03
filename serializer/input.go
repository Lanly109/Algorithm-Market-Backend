package serializer

import "singo/model"

// input 输入序列化器
type Input struct {
	ID    uint   `json:"id"`
	Input string `json:"input"`
}

// BuildInput 序列化输入
func BuildInputs(inputs []model.Input) []Input {
	var input []Input = []Input{}
	for _, i := range inputs {
		input = append(input, Input{
			ID:    i.ID,
			Input: i.Input,
		})
	}
	return input
}

// BuildInputResponse 序列化输入响应
func BuildInputsResponse(inputs []model.Input) Response {
	return Response{
		Data: BuildInputs(inputs),
	}
}

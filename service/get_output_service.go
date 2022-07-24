package service

import (
	"singo/model"
	"singo/serializer"
)

// GetInputListService 获得输入列表
type GetOutputService struct {
	ID  uint32 `uri:"id"`
	IID uint32 `uri:"inputID"`
}

// GetData 获取数据
func (service *GetOutputService) GetData() serializer.Response {

	input, _ := model.GetInput(service.IID)

	return serializer.BuildOutputResponse(input)
}

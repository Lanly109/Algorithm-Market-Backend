package service

import (
	"singo/model"
	"singo/serializer"
)

// GetInputListService 获得输入列表
type GetInputListService struct {
	ID uint32 `uri:"id"`
}

// GetData 获取数据
func (service *GetInputListService) GetData() serializer.Response {

    inputs, _ := model.GetInputs(service.ID)
	
	return serializer.BuildInputsResponse(inputs)
}

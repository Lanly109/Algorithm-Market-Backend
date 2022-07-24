package service

import (
	"singo/model"
	"singo/serializer"
)

// GetItemFullDeatilService 获得商品完整信息服务
type GetItemFullDetailService struct {
	ID uint `uri:"id"`
}

// GetData 获取数据
func (service *GetItemFullDetailService) GetData() serializer.Response {
	print(service.ID)
	item, err := model.GetItem(service.ID)

	if err != nil {
		return serializer.ParamErr("商品不存在！", nil)
	}

	tags, err := model.GetTags(service.ID)
	inputs, err := model.GetInputs(service.ID)

	return serializer.BuildItemFullResponse(item, tags, inputs)
}

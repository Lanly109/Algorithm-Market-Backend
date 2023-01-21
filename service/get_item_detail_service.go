package service

import (
	"singo/model"
	"singo/serializer"
)

// GetItemDeatilService 获得商品详细信息服务
type GetItemDetailService struct {
	ID uint `uri:"id"`
}

// GetData 获取数据
func (service *GetItemDetailService) GetData() serializer.Response {
	item, err := model.GetItem(service.ID)

	if err != nil {
		return serializer.ParamErr("商品不存在！", nil)
	}

	tags, err := model.GetTags(service.ID)
	inputs, err := model.GetInputs(service.ID)

	return serializer.BuildItemResponse(item, tags, inputs)
}

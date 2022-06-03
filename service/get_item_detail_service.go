package service

import (
	"singo/model"
	"singo/serializer"
)

// GetItemDeatilService 获得商品详细列表服务
type GetItemDetailService struct {
	ID uint `uri:"id"`
}

// GetData 获取数据
func (service *GetItemDetailService) GetData() serializer.Response {
    print(service.ID)
	item, err := model.GetItem(service.ID)

	if err != nil {
		return serializer.ParamErr("商品不存在！", nil)
	}

	tags, err := model.GetTags(service.ID)

	return serializer.BuildItemResponse(item, tags)
}

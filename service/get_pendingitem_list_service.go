package service

import (
	"singo/model"
	"singo/serializer"
)

// GetPendingItemListService 获得自己商品列表服务
type GetPendingItemListService struct {
}

// GetData 获取数据
func (service *GetPendingItemListService) GetData() serializer.Response {
	var items []model.Item

	model.DB.Where("status = ?", model.Pending).Find(&items)

	var tags [][]model.Tag

	for _, j := range items {
		tag, _ := model.GetTags(j.ID)
		tags = append(tags, tag)
	}

	return serializer.BuildItemListResponse(items, tags, true)
}

package service

import (
	"singo/model"
	"singo/serializer"
)

// GetMyItemListService 获得自己商品列表服务
type GetMyItemListService struct {
}

// GetData 获取数据
func (service *GetMyItemListService) GetData(user *model.User) serializer.Response {
	var items []model.Item

	model.DB.Where("author_id = ?", user.ID).Find(&items)

	var tags [][]model.Tag

	for _, j := range items {
		tag, _ := model.GetTags(j.ID)
		tags = append(tags, tag)
	}

	return serializer.BuildItemListResponse(items, tags, true)
}

package service

import (
	"singo/model"
	"singo/serializer"
)

// GetItemDeatilService 获得商品详细列表服务
type GetItemListService struct {
}

// GetData 获取数据
func (service *GetItemListService) GetData() serializer.Response {
    var items []model.Item

    model.DB.Find(&items)
	
    var tags [][]model.Tag

    for _, j := range items{
        tag, _ := model.GetTags(j.ID)
        tags = append(tags, tag)
    }

	return serializer.BuildItemListResponse(items, tags)
}

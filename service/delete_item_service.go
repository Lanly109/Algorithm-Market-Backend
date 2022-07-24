package service

import (
	"singo/model"
	"singo/serializer"
)

// DeleteItemService 删除商品服务
type DeleteItemService struct {
	ID uint `uri:"id"`
}

// DeleteData 删除数据
func (service *DeleteItemService) DeleteData() serializer.Response {

	if err := model.DeleteTags(service.ID); err != nil {
		return serializer.DBErr("删除标签出错", err)
	}

	if err := model.DeleteInputs(service.ID); err != nil {
		return serializer.DBErr("删除输入样例出错", err)
	}

	if err := model.DB.Delete(&model.Item{}, service.ID).Error; err != nil {
		return serializer.DBErr("删除商品出错", err)
	}

	return serializer.OK()
}

package service

import (
	"singo/model"
	"singo/serializer"

	"gorm.io/gorm"
)

// DeleteItemService 删除商品服务
type DeleteItemService struct {
	ID uint `uri:"id"`
}

// DeleteData 删除数据
func (service *DeleteItemService) DeleteData() serializer.Response {

	err := model.DB.Transaction(func(tx *gorm.DB) error {
		if err := model.DeleteTags(service.ID); err != nil {
			return err
		}

		if err := model.DeleteInputs(service.ID); err != nil {
			return err
		}

		if err := model.DB.Delete(&model.Item{}, service.ID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return serializer.ServiceErr("删除商品失败", err)
	}

	return serializer.OK()
}

package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"
)

// UpdateItemService 更新商品服务
type UpdateItemService struct {
	ID uint `uri:"id"`
	CreateItemService
}

// UpdateData 更新数据
func (service *UpdateItemService) UpdateData() serializer.Response {

	if err := model.DeleteTags(service.ID); err != nil {
		return serializer.DBErr("删除标签出错", err)
	}

	if err := model.DeleteInputs(service.ID); err != nil {
		return serializer.DBErr("删除输入样例出错", err)
	}

	if err := model.CreateTags(service.ID, service.Tag); err != nil {
		return serializer.ParamErr("创建标签失败", err)
	}

	if err := model.CreateInputs(service.ID, service.Input); err != nil {
		return serializer.ParamErr("创建输入数据失败", err)
	}

	fmt.Println(service)
	var item model.Item
	model.DB.Find(&item, service.ID)
	item.Name = service.Name
	item.Brief = service.Brief
	item.Picture = service.Picture
	item.Introduce = service.Introduce
	item.Algorithm = service.Algorithm
	item.Code = service.Code
	item.Time = service.Time
	item.Memory = service.Memory
	item.OutputImg = service.OutputImg
	model.DB.Save(&item)

	return serializer.OK()
}

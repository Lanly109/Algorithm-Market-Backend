package service

import (
    "singo/model"
    "singo/serializer"

    "gorm.io/gorm"
)

// UpdateItemService 更新商品服务
type UpdateItemService struct {
    ID uint `uri:"id"`
    CreateItemService
}

// UpdateData 更新数据
func (service *UpdateItemService) UpdateData() serializer.Response {

    err := model.DB.Transaction(func(tx *gorm.DB) error {
        if err := model.DeleteTags(service.ID); err != nil {
            return err
        }

        if err := model.DeleteInputs(service.ID); err != nil {
            return err
        }

        if err := model.CreateTags(service.ID, service.Tag); err != nil {
            return err
        }

        if err := model.CreateInputs(service.ID, service.Input); err != nil {
            return err
        }

        var item model.Item
        if err := model.DB.Find(&item, service.ID).Error; err != nil{
            return err
        }
        item.Name = service.Name
        item.Brief = service.Brief
        item.Picture = service.Picture
        item.Introduce = service.Introduce
        item.Algorithm = service.Algorithm
        item.Code = service.Code
        item.Time = service.Time
        item.Memory = service.Memory
        item.OutputImg = service.OutputImg
        item.Status = model.Pending
        if err := model.DB.Save(&item).Error; err != nil{
            return err
        }

        return nil
    })

    if err != nil {
        return serializer.ServiceErr("更新商品失败", err)
    }

    return serializer.OK()
}

package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"
)

// CreateItemService 创建商品服务
type CreateItemService struct {
	Name      string   `json:"name"`
	Brief     string   `json:"brief"`
	Picture   string   `json:"picture"`
	Tag       []string `json:"tag"`
	Input     []string `json:"input"`
	Introduce string   `json:"introduce"`
	Algorithm string   `json:"algorithm"`
	Code      string   `json:"code"`
}

// CreateData 创建数据
func (service *CreateItemService) CreateData() serializer.Response {

    fmt.Println(service)
	item := model.Item{
		Name:      service.Name,
		Brief:     service.Brief,
		Picture:   service.Picture,
		Introduce: service.Introduce,
		Algorithm: service.Algorithm,
		Code:      service.Code,
	}

	if err := model.DB.Create(&item).Error; err != nil {
		return serializer.ParamErr("创建商品失败", err)
	}

	if err := model.CreateTags(item.ID, service.Tag); err != nil {
		return serializer.ParamErr("创建标签失败", err)
	}

	if err := model.CreateInputs(item.ID, service.Input); err != nil {
		return serializer.ParamErr("创建输入数据失败", err)
	}

	return serializer.OK()
}

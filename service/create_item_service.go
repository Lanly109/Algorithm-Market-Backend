package service

import (
	"singo/model"
	"singo/serializer"

	"gorm.io/gorm"
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
	Time      int      `json:"time"`
	Memory    int      `json:"memory"`
	OutputImg bool     `json:"output_img"`
}

// CreateData 创建数据
func (service *CreateItemService) CreateData(user *model.User) serializer.Response {

	item := model.Item{
		Name:      service.Name,
		Brief:     service.Brief,
		Picture:   service.Picture,
		Introduce: service.Introduce,
		Algorithm: service.Algorithm,
		Code:      service.Code,
		Time:      service.Time,
		Memory:    service.Memory,
		OutputImg: service.OutputImg,
		Status:    model.Pending,
		AuthorId:  user.ID,
	}

	err := model.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&item).Error; err != nil {
			return err
		}

		if err := model.CreateTags(item.ID, service.Tag); err != nil {
			return err
		}

		if err := model.CreateInputs(item.ID, service.Input); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return serializer.ServiceErr("创建商品失败", err)
	}

	return serializer.OK()
}

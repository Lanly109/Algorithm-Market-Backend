package service

import (
	"singo/model"
	"singo/serializer"
	"time"
)

// AcceptItemService 审核通过商品
type AcceptItemService struct {
	ID uint `uri:"id"`
}

// UpdateData 更新数据
func (service *AcceptItemService) UpdateData(user *model.User) serializer.Response {
	item, err := model.GetItem(service.ID)

	if err != nil {
		return serializer.ParamErr("商品不存在！", nil)
	}

	item.Status = model.Accept
	item.ReviewId = user.ID
    now := time.Now()
	item.ReviewTime = &now
	model.DB.Save(&item)

	return serializer.OK()
}

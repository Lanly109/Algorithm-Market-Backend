package service

import (
	"singo/model"
	"singo/serializer"
	"time"
)

// RejectItemService 审核通过商品
type RejectItemService struct {
	ID uint `uri:"id"`
}

// UpdateData 更新数据
func (service *RejectItemService) UpdateData(user *model.User) serializer.Response {
	item, err := model.GetItem(service.ID)

	if err != nil {
		return serializer.ParamErr("商品不存在！", nil)
	}

	item.Status = model.Reject
	item.ReviewId = user.ID
    now := time.Now()
	item.ReviewTime = &now
	model.DB.Save(&item)

	return serializer.OK()
}

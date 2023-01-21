package api

import (
	"singo/middleware"
	"singo/model"
	"singo/service"

	"github.com/gin-gonic/gin"
)

// GetItemDetail 获取商品详细信息接口
func GetItemDetail(c *gin.Context) {
	if middleware.IsAdmin(c) {
		var service service.GetItemFullDetailService
		if err := c.ShouldBindUri(&service); err == nil {
			res := service.GetData()
			c.JSON(200, res)
		} else {
			c.JSON(200, ErrorResponse(err))
		}
	} else {
		var service service.GetItemDetailService
		if err := c.ShouldBindUri(&service); err == nil {
			res := service.GetData()
			c.JSON(200, res)
		} else {
			c.JSON(200, ErrorResponse(err))
		}
	}
}

// GetItemFullDetail 获取商品完整信息接口
func GetItemFullDetail(c *gin.Context) {
	var service service.GetItemDetailService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.GetData()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetItemList 获取商品列表接口
func GetItemList(c *gin.Context) {
	var service service.GetItemListService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.GetData()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetMyItemList 获取自己商品列表接口
func GetMyItemList(c *gin.Context) {
	var service service.GetMyItemListService
	if err := c.ShouldBindUri(&service); err == nil {
		user, _ := c.Get("user")
		real, _ := user.(*model.User)
		res := service.GetData(real)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetPendingItemList 获取待审核商品列表接口
func GetPendingItemList(c *gin.Context) {
	var service service.GetPendingItemListService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.GetData()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CreateItem 创建商品接口
func CreateItem(c *gin.Context) {
	var service service.CreateItemService
	if err := c.ShouldBind(&service); err == nil {
		user, _ := c.Get("user")
		real, _ := user.(*model.User)
		res := service.CreateData(real)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteItem 删除商品接口
func DeleteItem(c *gin.Context) {
	var service service.DeleteItemService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.DeleteData()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UpdateItem 更新商品接口
func UpdateItem(c *gin.Context) {
	var service service.UpdateItemService
	if err := c.ShouldBindUri(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	}
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateData()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AcceptItem 审核通过商品
func AcceptItem(c *gin.Context) {
	var service service.AcceptItemService
	if err := c.ShouldBindUri(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	}
	user, _ := c.Get("user")
	real, _ := user.(*model.User)
	res := service.UpdateData(real)
	c.JSON(200, res)
}

// RejectItem 审核拒绝商品
func RejectItem(c *gin.Context) {
	var service service.RejectItemService
	if err := c.ShouldBindUri(&service); err != nil {
		c.JSON(200, ErrorResponse(err))
	}
	user, _ := c.Get("user")
	real, _ := user.(*model.User)
	res := service.UpdateData(real)
	c.JSON(200, res)
}

package api

import (
	"singo/middleware"
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

// CreateItem 创建商品接口
func CreateItem(c *gin.Context) {
	var service service.CreateItemService
	if err := c.ShouldBind(&service); err == nil {
		res := service.CreateData()
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

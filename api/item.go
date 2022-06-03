package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

// GetItemDetail 获取商品详细信息接口
func GetItemDetail(c *gin.Context) {
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

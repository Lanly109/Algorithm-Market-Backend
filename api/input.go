package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

// GetInputList 获取输入列表
func GetInputList(c *gin.Context) {
	var service service.GetInputListService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.GetData()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetOutput 获取输出结果
func GetOutput(c *gin.Context) {
	var service service.GetOutputService
	if err := c.ShouldBindUri(&service); err == nil {
		res := service.GetData()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

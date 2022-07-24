package api

import (
	"singo/service"

	"github.com/gin-gonic/gin"
)

// Judge 评测接口
func Judge(c *gin.Context) {
	var service service.JudgeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Judge()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}


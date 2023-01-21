package api

import (
	"fmt"
	"path"
	"path/filepath"
	"singo/serializer"
	"singo/util"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ImagePath    = "image"
	ImageUrlPath = "/img/"
	LimitSize    = 2 * 1024 * 1024 // bytes
)

// UploadImage 图片上传接口
func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	var res serializer.Response
	if err != nil {
		res = serializer.ServiceErr("获取图片失败", err)
	} else {
		fileExt := strings.ToLower(path.Ext(file.Filename))
		fmt.Println(file.Size)
		if file.Size > LimitSize || (fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg") {
			res = serializer.ParamErr("仅支持小于2M的png、jpg、jpeg格式的图片", nil)
		} else {
			name := util.Md5(fmt.Sprintf("%s%s", file.Filename, time.Now().String())) + fileExt
			dest := filepath.Join(ImagePath, name)
			err = c.SaveUploadedFile(file, dest)
            if err != nil{
                res = serializer.ServiceErr("保存图片出错", err)
            }else{
                res = serializer.OK()
                res.Data = ImageUrlPath + name
            }
		}
	}

	c.JSON(200, res)
}

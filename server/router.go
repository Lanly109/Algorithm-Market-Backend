package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api")
	{
		v1.POST("ping", api.Ping)

        // 文章列表
		v1.GET("items", api.GetItemList)

        // 审核通过文章
		accept := v1.Group("")
        accept.Use(middleware.AcceptItemRequired())
        {
            // 获取商品详细信息
            accept.GET("items/:id", api.GetItemDetail)
        }

        // 评测
		v1.POST("judge", api.Judge)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 图片上传
		v1.POST("img", api.UploadImage)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// 用户信息
			auth.GET("user/info", api.UserMe)
			// 用户登出
			auth.DELETE("user/logout", api.UserLogout)

            // 创建商品
            auth.POST("items", api.CreateItem)

            auth.GET("myitems", api.GetMyItemList)

			// 需要所有者或管理员权限的
			owner := auth.Group("")
			owner.Use(middleware.OwnerRequired())
			{
				owner.PUT("items/:id", api.UpdateItem)
				owner.DELETE("items/:id", api.DeleteItem)
			}

            admin := auth.Group("")
            admin.Use(middleware.AdminRequired())
            {
                admin.GET("items/:id/accept", api.AcceptItem)
                admin.GET("items/:id/reject", api.RejectItem)
                admin.GET("reviewitems", api.GetPendingItemList)
            }
		}

	}
	return r
}

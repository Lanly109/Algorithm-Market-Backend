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

		v1.GET("items", api.GetItemList)
		v1.GET("items/:id", api.GetItemDetail)

		v1.POST("judge", api.Judge)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/info", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)

			// 需要管理员权限的
			admin := auth.Group("")
			admin.Use(middleware.AdminRequired())
			{
				// User Routing
				admin.POST("items", api.CreateItem)
				admin.PUT("items/:id", api.UpdateItem)
				admin.DELETE("items/:id", api.DeleteItem)
			}
		}

	}
	return r
}

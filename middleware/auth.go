package middleware

import (
	"singo/model"
	"singo/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}

// AdminRequired 需要管理员
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if IsAdmin(c) {
			c.Next()
			return
		}
		c.JSON(200, serializer.PermissionDeny())
		c.Abort()
	}
}

// AcceptItemRequired 通过审核的文章
func AcceptItemRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if IsOwner(c) || IsAdmin(c) {
			c.Next()
			return
		} else if item, err := model.GetItem(id); err == nil && item.Status == model.Accept {
			c.Next()
			return
		}
		c.JSON(200, serializer.PermissionDeny())
		c.Abort()
	}
}

// OwnerRequired 需要所有者或管理员
func OwnerRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if IsOwner(c) || IsAdmin(c) {
			c.Next()
			return
		}
		c.JSON(200, serializer.PermissionDeny())
		c.Abort()
	}
}

func IsAdmin(c *gin.Context) bool {
	if user, _ := c.Get("user"); user != nil {
		if real, ok := user.(*model.User); ok {
			if real.IsAdmin() {
				return true
			}
		}
	}
	return false
}

func IsOwner(c *gin.Context) bool {
	if user, _ := c.Get("user"); user != nil {
		id := c.Param("id")
		if item, err := model.GetItem(id); err == nil {
			if real, ok := user.(*model.User); ok {
				if real.ID == uint(item.AuthorId) {
					return true
				}
			}
		}
	}
	return false
}

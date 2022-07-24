package serializer

import "singo/model"

// User 用户序列化器
type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Coin     int    `json:"coin"`
	Avatar   string `json:"avatar"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
		Role:     user.Role,
		Coin:     user.Coin,
		Avatar:   user.Avatar,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}

package api

type CreateUserRequest struct {
	Username        string `json:"username" name:"用户名" validate:"required"`
	Password        string `json:"password" name:"密码" validate:"required"`
	ConfirmPassword string `json:"confirm_password" name:"确认密码" validate:"required,eqfield=Password"`
}

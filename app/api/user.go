package api

type CreateUserRequest struct {
	Username        string   `json:"username" name:"用户名" validate:"required"`
	Password        string   `json:"password" name:"密码" validate:"required"`
	ConfirmPassword string   `json:"confirm_password" name:"确认密码" validate:"required,eqfield=Password" message:"eqfield:两次密码输入不一致"`
	RoleIds         []string `json:"role_ids"`
}

type UserListRequest struct {
	Current  int    `form:"current"`
	PageSize int    `form:"pageSize"`
	Username string `form:"username"`
}
type UserListItem struct {
	Id       uint     `json:"id"`
	Username string   `json:"username"`
	Role     []string `json:"role" gorm:"-"`
}
type UserListResponse struct {
	Total int64          `json:"total"`
	Data  []UserListItem `json:"data"`
}

type GetUserRequest struct {
	Id uint `form:"id" validate:"required"`
}
type GetUserResponse struct {
	Id       uint     `json:"id"`
	Username string   `json:"username"`
	RoleIds  []string `json:"role_ids" gorm:"-"`
}

type EditUserRequest struct {
	Id              uint     `json:"id" validate:"required"`
	Username        string   `json:"username" name:"用户名" validate:"required"`
	Password        string   `json:"password" name:"密码"`
	ConfirmPassword string   `json:"confirm_password" name:"确认密码" validate:"required_with=Password,eqfield=Password" message:"required_with:请确认密码,eqfield:两次密码输入不一致"`
	RoleIds         []string `json:"role_ids"`
}

type UserSelectListResponse struct {
	Label string `json:"label" gorm:"column:username"`
	Value string `json:"value" gorm:"column:id"`
}

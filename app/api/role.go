package api

type CreateRoleRequest struct {
	Name    string `json:"name" name:"名称" validate:"required"`
	MenuIds []int  `json:"menu_ids"`
}

type RoleListRequest struct {
	Current  int    `form:"current"`
	PageSize int    `form:"pageSize"`
	Name     string `form:"name"`
}
type RoleListItem struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
type RoleListResponse struct {
	Total int64          `json:"total"`
	Data  []RoleListItem `json:"data"`
}

type GetRoleRequest struct {
	Id uint `form:"id" validate:"required"`
}
type GetRoleResponse struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	MenuIds []int  `json:"menu_ids"`
}

type EditRoleRequest struct {
	Id uint `json:"id"`
	CreateRoleRequest
}

type UserRoleListResponse struct {
	Label string `json:"label" gorm:"column:name"`
	Value string `json:"value" gorm:"column:id"`
}

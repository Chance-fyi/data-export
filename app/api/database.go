package api

type DatabaseListRequest struct {
	Current  int    `form:"current"`
	PageSize int    `form:"pageSize"`
	Name     string `form:"name"`
}
type DatabaseListItem struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
type DatabaseListResponse struct {
	Total int64              `json:"total"`
	Data  []DatabaseListItem `json:"data"`
}

type CreateDatabaseRequest struct {
	Name     string `json:"name" name:"名称" validate:"required"`
	Hostname string `json:"hostname" name:"数据库地址" validate:"required"`
	Port     string `json:"port" name:"端口" validate:"required"`
	Username string `json:"username" name:"用户名" validate:"required"`
	Password string `json:"password" name:"密码" validate:"required"`
	Database string `json:"database" name:"数据库名" validate:"required"`
}

type GetDatabaseRequest struct {
	Id int `form:"id"`
}
type GetDatabaseResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type EditDatabaseRequest struct {
	Id uint `json:"id" validate:"required"`
	GetDatabaseResponse
}

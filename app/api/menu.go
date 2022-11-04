package api

type CreateMenuRequest struct {
	Name     string `json:"name" name:"名称" validate:"required"`
	Path     string `json:"path" name:"Path"`
	ParentId uint   `json:"parent_id"`
}

type EditMenuRequest struct {
	Id uint `json:"id" validate:"required"`
	CreateMenuRequest
}

type GetMenuRequest struct {
	Id uint `form:"id" validate:"required"`
}
type GetMenuResponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	ParentId uint   `json:"parent_id"`
}

type MenuListRequest struct {
	Current  int    `form:"current"`
	PageSize int    `form:"pageSize"`
	Name     string `form:"name"`
}
type MenuListItem struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}
type MenuListResponse struct {
	Total int64          `json:"total"`
	Data  []MenuListItem `json:"data"`
}

type MenuSelectTreeResponse struct {
	Value    uint                     `json:"value"`
	Title    string                   `json:"title"`
	Children []MenuSelectTreeResponse `json:"children"`
}

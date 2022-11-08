package api

type CreateSqlRequest struct {
	Sql string `json:"sql" name:"SQL" validate:"required"`
}

type SqlListRequest struct {
	Current  int    `form:"current"`
	PageSize int    `form:"pageSize"`
	Fields   string `form:"fields"`
}
type SqlListItem struct {
	Id     uint   `json:"id"`
	Fields string `json:"fields"`
}
type SqlListResponse struct {
	Total int64         `json:"total"`
	Data  []SqlListItem `json:"data"`
}

type GetSqlRequest struct {
	Id int `form:"id"`
}
type GetSqlResponse struct {
	Id  int    `json:"id"`
	Sql string `json:"sql"`
}

type EditSqlRequest struct {
	Id uint `json:"id" validate:"required"`
	CreateSqlRequest
}

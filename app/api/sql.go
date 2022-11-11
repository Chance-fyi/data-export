package api

type CreateSqlRequest struct {
	DatabaseId string `json:"database_id" name:"数据库" validate:"required"`
	Name       string `json:"name" name:"备注" validate:"required"`
	Sql        string `json:"sql" name:"SQL" validate:"required"`
}

type SqlListRequest struct {
	Current    int    `form:"current"`
	PageSize   int    `form:"pageSize"`
	Fields     string `form:"fields"`
	DatabaseId string `form:"database_id"`
}
type SqlListItem struct {
	Id           uint   `json:"id"`
	Fields       string `json:"fields"`
	Name         string `json:"name"`
	DatabaseName string `json:"database_name"`
}
type SqlListResponse struct {
	Total int64         `json:"total"`
	Data  []SqlListItem `json:"data"`
}

type GetSqlRequest struct {
	Id int `form:"id"`
}
type GetSqlResponse struct {
	Id         uint   `json:"id"`
	Sql        string `json:"sql"`
	Name       string `json:"name"`
	DatabaseId string `json:"database_id"`
}

type EditSqlRequest struct {
	Id uint `json:"id" validate:"required"`
	CreateSqlRequest
}

type GetUserSqlRequest struct {
	Id int `form:"id" validate:"required"`
}
type GetUserSqlResponse struct {
	Id      int      `json:"id"`
	UserIds []string `json:"user_ids"`
}

type SetUserSqlRequest struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	UserIds []string `json:"user_ids"`
}

type MySqlListRequest struct {
	Current    int    `form:"current"`
	PageSize   int    `form:"pageSize"`
	Name       string `form:"name"`
	Fields     string `form:"fields"`
	DatabaseId string `form:"database_id"`
}
type MySqlListItem struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Fields       string `json:"fields"`
	SqlId        uint   `json:"sql_id"`
	DatabaseName string `json:"database_name"`
}
type MySqlListResponse struct {
	Total int64           `json:"total"`
	Data  []MySqlListItem `json:"data"`
}

type GetUserSqlNameRequest struct {
	Id int `form:"id" validate:"required"`
}
type GetUserSqlNameResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type SetUserSqlNameRequest struct {
	Id   uint   `json:"id" validate:"required"`
	Name string `json:"name" name:"备注" validate:"required"`
}

type GetDownloadSqlRequest struct {
	Id int `form:"id" validate:"required"`
}
type GetDownloadSqlField struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
type GetDownloadSqlResponse struct {
	Id     uint                  `json:"id"`
	Sql    string                `json:"sql"`
	Fields []GetDownloadSqlField `json:"fields"`
}

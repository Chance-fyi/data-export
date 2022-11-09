package api

type CreateExportRequest struct {
	Id       uint   `json:"id" validate:"required"`
	Filename string `json:"filename" name:"文件名" validate:"required"`
	Sql      string `json:"sql" name:"SQL" validate:"required"`
}

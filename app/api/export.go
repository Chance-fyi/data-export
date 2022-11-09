package api

type CreateExportRequest struct {
	Id       uint   `json:"id" validate:"required"`
	Filename string `json:"filename" name:"文件名" validate:"required"`
	Sql      string `json:"sql" name:"SQL" validate:"required"`
}

type ExportListRequest struct {
	Current  int    `form:"current"`
	PageSize int    `form:"pageSize"`
	Filename string `form:"filename"`
	Status   string `form:"status"`
}
type ExportListItem struct {
	Id       uint   `json:"id"`
	Filename string `json:"filename"`
	STATUS   uint   `json:"status"`
	ErrorMsg string `json:"error_msg"`
}
type ExportListResponse struct {
	Total int64            `json:"total"`
	Data  []ExportListItem `json:"data"`
}

type ExportDownloadRequest struct {
	Id uint `form:"id"`
}

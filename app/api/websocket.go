package api

type WsResponse struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type WsExportResponse struct {
	Name string `json:"name"`
}

package api

type LoginRequest struct {
	Username string `json:"username" name:"用户名" validate:"required"`
	Password string `json:"password" name:"密码" validate:"required"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

type RefreshTokenRequest struct {
	Token string `json:"token" validate:"required"`
}
type RefreshTokenResponse struct {
	Token string `json:"token"`
}

type UsesMenuList struct {
	Name   string         `json:"name"`
	Path   string         `json:"path"`
	Routes []UsesMenuList `json:"routes"`
}

type GetUserInfoResponse struct {
	Id       uint           `json:"id"`
	Username string         `json:"username"`
	Menu     []UsesMenuList `json:"menu"`
}

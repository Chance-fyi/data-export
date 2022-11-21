package controller

import (
	"data-export/pkg/jwt"
	ws "data-export/pkg/websocket"
	"github.com/tidwall/gjson"
)

func WsLogin(json gjson.Result, c *ws.Client) error {
	token := json.Get("token").String()
	claims, err := jwt.Parse(token)
	if err != nil {
		return err
	}
	ws.SetTag(claims.ID, c)
	return nil
}

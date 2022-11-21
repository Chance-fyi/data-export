package boot

import (
	"data-export/app/controller"
	ws "data-export/pkg/websocket"
	"github.com/tidwall/gjson"
)

func initWebSocket() {
	ws.CreateServer(map[string]func(json gjson.Result, c *ws.Client) error{
		"login": controller.WsLogin,
	})
}

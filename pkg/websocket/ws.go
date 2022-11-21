package ws

import (
	"data-export/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	conn *websocket.Conn
}

type server struct {
	clients    map[*Client]string
	tags       map[string]*Client
	register   chan *Client
	unregister chan *Client
	fun        map[string]func(json gjson.Result, c *Client) error
}

var (
	once     sync.Once
	s        *server
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func CreateServer(fun map[string]func(json gjson.Result, c *Client) error) {
	once.Do(func() {
		s = &server{
			clients:    make(map[*Client]string),
			tags:       make(map[string]*Client),
			register:   make(chan *Client),
			unregister: make(chan *Client),
			fun:        fun,
		}
		go s.run()
	})
}

func (*server) run() {
	for {
		select {
		case c := <-s.register:
			s.clients[c] = ""
		case c := <-s.unregister:
			delete(s.tags, s.clients[c])
			delete(s.clients, c)
		}
	}
}

func SetTag(tag string, c *Client) {
	delete(s.tags, s.clients[c])
	s.clients[c] = tag
	s.tags[tag] = c
}

func SendUseTag(tag string, message interface{}) {
	c, ok := s.tags[tag]
	if ok {
		_ = c.conn.WriteJSON(message)
	}
}

func ServeWs(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	c := &Client{conn: conn}
	s.register <- c
	go c.run()
}

func (c *Client) run() {
	defer func() {
		s.unregister <- c
		_ = c.conn.Close()
	}()
	_ = c.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
	for {
		_, ms, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		if string(ms) == "ping" {
			_ = c.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
			continue
		}
		json := gjson.Parse(string(ms))
		f, ok := s.fun[json.Get("type").String()]
		if !ok {
			break
		}
		err = f(json.Get("data"), c)
		if err != nil {
			break
		}
	}
}

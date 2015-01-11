package controllers

import (
	xgo "github.com/nexigolabs/nexigo"
	"net/http"
)

type WebsocketController struct {
	xgo.Controller
}

func (c *WebsocketController) Index(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Websocket"
	c.ServeTpl([]string{"views/app/websocket.html", "views/header.html", "views/footer.html"}, data)
}

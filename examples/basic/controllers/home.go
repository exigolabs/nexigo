package controllers

import (
	xgo "github.com/nexigolabs/nexigo"
	"net/http"
)

type HomeController struct {
	xgo.Controller
}

func (c *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Nexigo"
	c.ServeTpl([]string{"views/app/home.html", "views/header.html", "views/footer.html"}, data)
}

func (c *HomeController) Public(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["public"] = xgo.PublicPath
	data["port"] = xgo.DefaultPort
	c.ServeJson(data)
}

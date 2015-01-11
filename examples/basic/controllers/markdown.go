package controllers

import (
	xgo "github.com/nexigolabs/nexigo"
	"net/http"
)

type MarkdownController struct {
	xgo.Controller
}

func (c *MarkdownController) Index(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Markdown"
	c.ServeTpl([]string{"views/app/markdown.html", "views/header.html", "views/footer.html"}, data)
}

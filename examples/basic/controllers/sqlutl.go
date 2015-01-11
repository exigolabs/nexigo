package controllers

import (
	xgo "github.com/nexigolabs/nexigo"
	"net/http"
)

type SqlController struct {
	xgo.Controller
}

func (c *SqlController) Index(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Sql Util"
	c.ServeTpl([]string{"views/app/sqlutl.html", "views/header.html", "views/footer.html"}, data)
}

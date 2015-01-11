package controllers

import (
	xgo "github.com/nexigolabs/nexigo"
	"net/http"
)

type CrudController struct {
	xgo.Controller
}

func (c *CrudController) Index(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Crud Sample"
	c.ServeTpl([]string{"views/app/crud.html", "views/header.html", "views/footer.html"}, data)
}

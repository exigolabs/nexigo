package controllers

import (
	"fmt"
	xgo "github.com/nexigolabs/nexigo"
	"net/http"
)

type ObjectController struct {
	xgo.Controller
}

func (c *ObjectController) Index(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Controller"
	c.ServeTpl([]string{"views/app/object.html", "views/header.html", "views/footer.html"}, data)
}

func (c *ObjectController) Text(w http.ResponseWriter, r *http.Request) {
	c.ServeText("<h1>Text</h1>Just demo <i style='color:blue'>sample</i> Text <br/><br/><a href='/object'>Back</a>")
}

func (c *ObjectController) Html(w http.ResponseWriter, r *http.Request) {
	c.ServeHtml("<h1>Html</h1>Just demo <i style='color:blue'>sample</i> Html <br/><br/><a href='/object'>Back</a>")
}

func (c *ObjectController) Template(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Template"
	c.ServeTpl([]string{"views/app/object/demo.html"}, data)
}

func (c *ObjectController) Json(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Json"
	data["FistName"] = "John"
	data["LastName"] = "Doe"
	c.ServeJson(data)
}

func (c *ObjectController) Form(w http.ResponseWriter, r *http.Request) {
	c.ServeTpl([]string{"views/app/object/form.html"}, nil)
}

func (c *ObjectController) Result(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	result := make(map[string]interface{})
	for k, v := range r.PostForm {
		result[k] = v[0]
	}
	c.ServeJson(result)
}

func test() {
	data := make(map[string]interface{})
	data["Title"] = "Template"
	fmt.Println("...", http.StateClosed)
}

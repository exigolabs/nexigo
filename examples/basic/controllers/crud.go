package controllers

import (
	"fmt"
	xgo "github.com/nexigolabs/nexigo"
	// "math/rand"
	"net/http"
	"strconv"
)

type CrudController struct {
	xgo.Controller
}

func (c *CrudController) Index(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["Title"] = "Crud Sample"
	c.ServeTpl([]string{"views/app/crud.html", "views/header.html", "views/footer.html"}, data)
}

func (c *CrudController) Insert(w http.ResponseWriter, r *http.Request) {
	sql := xgo.OpenDB()
	defer sql.CloseDB()

	sql.Execute("delete table1")
	for i := 1; i <= 100; i++ {
		text := "Text " + strconv.Itoa(i)
		sql.Execute("insert into table1(id, text) values(?,?)", i, text)
	}

	result := make(map[string]interface{})
	result["success"] = true
	result["data"] = sql.QueryToList("select top 10 * from table1 order by id")
	result["form"] = xgo.GetFormValue(r)
	c.ServeJson(result)
}

func (c *CrudController) Read(w http.ResponseWriter, r *http.Request) {
	sql := xgo.OpenDB()
	defer sql.CloseDB()

	list := sql.QueryToList("select top 15 * from table1 where id > ? order by id desc", "20")
	c.ServeJson(list)
}

func Text() {
	fmt.Println("...")
}

package main

import (
	"fmt"
	xgo "github.com/nexigolabs/nexigo"
	ctl "github.com/nexigolabs/nexigo/examples/basic/controllers"
)

func init() {
	fmt.Println("example - main")
}

func main() {
	xgo.Route("/", &ctl.HomeController{})
	xgo.Route("/object", &ctl.ObjectController{})
	xgo.Route("/crud/", &ctl.CrudController{})
	xgo.Route("/markdown", &ctl.MarkdownController{})
	xgo.Route("/websocket", &ctl.WebsocketController{})
	xgo.Route("/sql", &ctl.SqlController{})

	xgo.Run()
}

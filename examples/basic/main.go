package main

import (
	xgo "github.com/nexigolabs/nexigo"
	ctl "github.com/nexigolabs/nexigo/examples/basic/controllers"
)

func init() {
	xgo.SetConfig("public", "public")
	xgo.SetConfig("port", "8080")
}

func main() {
	// xgo.Route("/", &ctl.HomeController{})
	xgo.Route("/object", &ctl.ObjectController{})
	xgo.Route("/crud/", &ctl.CrudController{})
	xgo.Route("/markdown", &ctl.MarkdownController{})
	xgo.Route("/websocket", &ctl.WebsocketController{})
	xgo.Route("/sql", &ctl.SqlController{})

	xgo.Run()
}

package nexigo

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

var (
	DefaultPort = "8080"
	PublicPath  = "public"
	HttpContext = ContextHandler{}
	Config      = make(map[string]interface{})
)

type RouteHandler struct {
	Path        string
	IController interface{}
}

type ContextHandler struct {
	routes []RouteHandler
}

func Route(path string, ctl interface{}) {
	HttpContext.routes = append(HttpContext.routes, RouteHandler{path, ctl})
}

func Run() {
	routes := HttpContext.routes

	http.Handle("/", http.FileServer(http.Dir(PublicPath)))

	for i := 0; i < len(routes); i++ {
		route := routes[i]
		path := strings.TrimSpace(route.Path)

		if path[:1] != "/" {
			path = "/" + path
		}

		if path[len(path)-1:len(path)] != "/" {
			path = path + "/"
		}

		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			rqpath := strings.TrimSpace(r.URL.Path)
			action := strings.ToLower(strings.Split(rqpath[len(path):], "/")[0])

			if action == "" {
				action = "index"
			}

			if action != "favicon.ico" {
				typeCont := reflect.TypeOf(route.IController)
				for i := 0; i < typeCont.NumMethod(); i++ {
					method := strings.ToLower(typeCont.Method(i).Name)
					if method == action {
						// reflect.ValueOf(route.IController).Method(i).Call([]reflect.Value{})

						fmt.Println("form ori", r)

						reflect.ValueOf(route.IController).Method(i).Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)})
						ctrlObj := reflect.ValueOf(route.IController).Elem().Field(0).Interface().(Controller)
						ctrlObj.RunAction(w, r)
						break
					}
				}
			} else {
				// fmt.Println("... favicon ...")
			}

		})
	}

	http.ListenAndServe(":"+DefaultPort, nil)
}

func SetConfig(name string, value interface{}) {
	Config[name] = value

	switch name {
	case "public":
		PublicPath = value.(string)
	case "port":
		DefaultPort = value.(string)
	}
}

func GetConfig(name string) interface{} {
	return Config[name]
}

func Text() {
	text := " /demo test/a ska     "
	fmt.Println("...", reflect.TypeOf(HttpContext), strings.TrimSpace(text))
}

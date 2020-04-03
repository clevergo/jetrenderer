package main

import (
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/CloudyKit/jet/v3"
	"github.com/clevergo/clevergo"
	"github.com/clevergo/jetrender"
)

func main() {
	set := jet.NewHTMLSet("./views")
	set.SetDevelopmentMode(true) // debug
	renderer := jetrender.New(set)
	renderer.SetBeforeRender(func(w io.Writer, name string, vars jet.VarMap, data interface{}, ctx *clevergo.Context) error {
		vars.SetFunc("title", jet.Func(func(args jet.Arguments) reflect.Value {
			args.RequireNumOfArguments("title", 1, 1)
			return reflect.ValueOf(strings.Title(args.Get(0).String()))
		}))

		return nil
	})

	router := clevergo.NewRouter()
	router.Renderer = renderer
	router.Get("/", func(ctx *clevergo.Context) error {
		return ctx.Render(200, "index.tmpl", map[string]interface{}{
			"message": "hello world",
		})
	})
	http.ListenAndServe(":8080", router)
}

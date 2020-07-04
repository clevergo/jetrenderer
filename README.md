# Renderer implementation of Jet template engine
[![Build Status](https://travis-ci.org/clevergo/jetrenderer.svg?branch=master)](https://travis-ci.org/clevergo/jetrenderer)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/jetrenderer/badge.svg?branch=master)](https://coveralls.io/github/clevergo/jetrenderer?branch=master)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/clevergo.tech/jetrenderer?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/jetrenderer)](https://goreportcard.com/report/github.com/clevergo/jetrenderer)
[![Release](https://img.shields.io/github/release/clevergo/jetrenderer.svg?style=flat-square)](https://github.com/clevergo/jetrenderer/releases)

## Usage

See [example](https://github.com/clevergo/examples/tree/master/jetrenderer).

```go
package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v4"
	"clevergo.tech/clevergo"
	"clevergo.tech/jetrenderer"
)

func main() {
	renderer := jetrenderer.New(jet.NewHTMLSet("./views"))
	renderer.SetBeforeRender(func(w io.Writer, name string, vars jet.VarMap, data interface{}, ctx *clevergo.Context) error {
        // populate variables and functions
        // vars.Set("user", user) // user instance.
        // vars.SetFunc("T", translate) // i18n translator.
		return nil
	})

	app := clevergo.New()
	app.Renderer = renderer
	app.Get("/", func(ctx *clevergo.Context) error {
		return ctx.Render(http.StatusOK, "index.tmpl", map[string]interface{}{
			"message": "hello world",
		})
	})
	app.Run(":8080")
}
```
# Render implementation for Jet template engine
[![Build Status](https://travis-ci.org/clevergo/jetrender.svg?branch=master)](https://travis-ci.org/clevergo/jetrender)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/jetrender/badge.svg?branch=master)](https://coveralls.io/github/clevergo/jetrender?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/jetrender)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/jetrender)](https://goreportcard.com/report/github.com/clevergo/jetrender)
[![Release](https://img.shields.io/github/release/clevergo/jetrender.svg?style=flat-square)](https://github.com/clevergo/jetrender/releases)

## Usage

See [example](example).

```go
package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v3"
	"github.com/clevergo/clevergo"
	"github.com/clevergo/jetrender"
)

func main() {
	renderer := jetrender.New(jet.NewHTMLSet("./views"))
	renderer.SetBeforeRender(func(w io.Writer, name string, vars jet.VarMap, data interface{}, ctx *clevergo.Context) error {
        // populate variables and functions
        // vars.Set("user", user) // user instance.
        // vars.SetFunc("T", translate) // i18n translator.
		return nil
	})

	router := clevergo.NewRouter()
	router.Renderer = renderer
	router.Get("/", func(ctx *clevergo.Context) error {
		return ctx.Render(http.StatusOK, "index.tmpl", map[string]interface{}{
			"message": "hello world",
		})
	})
	http.ListenAndServe(":8080", router)
}
```
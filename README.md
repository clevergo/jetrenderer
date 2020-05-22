# Render implementation for Jet template engine
[![Build Status](https://travis-ci.org/clevergo/jetrenderer.svg?branch=master)](https://travis-ci.org/clevergo/jetrenderer)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/jetrenderer/badge.svg?branch=master)](https://coveralls.io/github/clevergo/jetrenderer?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/jetrenderer)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/jetrenderer)](https://goreportcard.com/report/github.com/clevergo/jetrenderer)
[![Sourcegraph](https://sourcegraph.com/github.com/clevergo/jetrenderer/-/badge.svg)](https://sourcegraph.com/github.com/clevergo/jetrenderer?badge)
[![Release](https://img.shields.io/github/release/clevergo/jetrenderer.svg?style=flat-square)](https://github.com/clevergo/jetrenderer/releases)

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
	"github.com/clevergo/jetrenderer"
)

func main() {
	renderer := jetrenderer.New(jet.NewHTMLSet("./views"))
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
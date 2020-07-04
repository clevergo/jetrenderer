// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package jetrenderer

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"

	"clevergo.tech/clevergo"
	"github.com/CloudyKit/jet/v4"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	set := jet.NewHTMLSet("")
	r := New(set)
	assert.Equal(t, set, r.Set)
}

func fakeBeforeRender(callback func(), err error) BeforeRender {
	return func(w io.Writer, name string, vars jet.VarMap, data interface{}, ctx *clevergo.Context) error {
		callback()
		return err
	}
}

func TestRenderer_SetBeforeRender(t *testing.T) {
	r := New(jet.NewHTMLSet(""))
	fn := fakeBeforeRender(func() {}, errors.New("before render test"))
	r.SetBeforeRender(fn)
	expected := fn(nil, "", nil, nil, nil)
	actual := r.beforeRender(nil, "", nil, nil, nil)
	assert.Equal(t, expected, actual)
}

func TestRenderer_Render(t *testing.T) {
	r := New(jet.NewHTMLSet("./testdata/views"))
	// invalid template.
	err := r.Render(nil, "404.tmpl", nil, nil)
	assert.NotNil(t, err)

	terminatedErr := errors.New("terminated")
	r.SetBeforeRender(fakeBeforeRender(func() {}, terminatedErr))
	err = r.Render(nil, "index.tmpl", nil, nil)
	assert.Equal(t, err, terminatedErr)

	r.SetBeforeRender(func(w io.Writer, name string, vars jet.VarMap, data interface{}, ctx *clevergo.Context) error {
		vars.SetFunc("title", jet.Func(func(args jet.Arguments) reflect.Value {
			args.RequireNumOfArguments("title", 1, 1)
			return reflect.ValueOf(strings.Title(args.Get(0).String()))
		}))

		return nil
	})
	for _, msg := range []string{"foo bar", "hello world"} {
		buf := new(bytes.Buffer)
		err = r.Render(buf, "index.tmpl", map[string]interface{}{
			"message": msg,
		}, nil)
		assert.Contains(t, buf.String(), strings.Title(msg))
	}
}

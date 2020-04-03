// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package jetrenderer

import (
	"io"
	"sync"

	"github.com/CloudyKit/jet/v3"
	"github.com/clevergo/clevergo"
)

var varsPool sync.Pool

func init() {
	varsPool = sync.Pool{
		New: func() interface{} {
			return jet.VarMap{}
		},
	}
}

func getVars() (vars jet.VarMap) {
	vars, _ = varsPool.Get().(jet.VarMap)
	for key := range vars {
		delete(vars, key)
	}
	return
}

func putVars(vars jet.VarMap) {
	varsPool.Put(vars)
}

// BeforeRender is a listener that listen the event which trigger before rendering.
type BeforeRender func(w io.Writer, name string, vars jet.VarMap, data interface{}, ctx *clevergo.Context) error

// Renderer implementation of clevergo.Renderer for Jet template engine.
type Renderer struct {
	*jet.Set
	beforeRender BeforeRender
}

// New returns a renderer with the given set.
func New(set *jet.Set) *Renderer {
	return &Renderer{
		Set: set,
	}
}

// SetBeforeRender sets before render listener.
func (r *Renderer) SetBeforeRender(fn BeforeRender) {
	r.beforeRender = fn
}

// Render implements clevergo.Renderer.Render.
func (r *Renderer) Render(w io.Writer, name string, data interface{}, ctx *clevergo.Context) error {
	t, err := r.GetTemplate(name)
	if err != nil {
		return err
	}

	vars := getVars()
	defer func() {
		putVars(vars)
	}()
	if r.beforeRender != nil {
		if err = r.beforeRender(w, name, vars, data, ctx); err != nil {
			return err
		}
	}

	return t.Execute(w, vars, data)
}

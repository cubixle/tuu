package tuu

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/gobuffalo/buffalo/render"
)

func NewContext(r Route, res http.ResponseWriter, req *http.Request) *DefaultContext {
	data := make(map[string]interface{})
	data["path"] = r.Path

	return &DefaultContext{
		response: res,
		request:  req,
		params:   req.URL.Query(),
		data:     data,
	}
}

type DefaultContext struct {
	context.Context
	response    http.ResponseWriter
	request     *http.Request
	params      url.Values
	contentType string
	data        map[string]interface{}
}

// Response returns the original Response for the request.
func (d *DefaultContext) Response() http.ResponseWriter {
	return d.response
}

// Request returns the original Request.
func (d *DefaultContext) Request() *http.Request {
	return d.request
}

// Params returns all of the parameters for the request,
// including both named params and query string parameters.
func (d *DefaultContext) Params() url.Values {
	return d.params
}

// Param returns a param, either named or query string,
// based on the key.
func (d *DefaultContext) Param(key string) string {
	return d.Params().Get(key)
}

// Set a value onto the Context. Any value set onto the Context
// will be automatically available in templates.
func (d *DefaultContext) Set(key string, value interface{}) {
	d.data[key] = value
}

// Value that has previously stored on the context.
func (d *DefaultContext) Value(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		if v, ok := d.data[k]; ok {
			return v
		}
	}
	return d.Context.Value(key)
}

func (d *DefaultContext) Render(status int, rr render.Renderer) error {
	if rr != nil {
		data := d.data
		pp := map[string]string{}
		for k, v := range d.params {
			pp[k] = v[0]
		}

		data["params"] = pp
		data["request"] = d.Request()
		bb := &bytes.Buffer{}

		err := rr.Render(bb, data)
		if err != nil {
			return err
		}

		d.Response().Header().Set("Content-Type", rr.ContentType())
		d.Response().WriteHeader(status)
		_, err = io.Copy(d.Response(), bb)
		if err != nil {
			return err
		}

		return nil
	}

	d.Response().WriteHeader(status)
	return nil
}

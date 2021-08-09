package context

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/worbridg/clean-architecture/user"
)

func NewContext(v ...interface{}) user.Context {
	if len(v) == 2 {
		return &HTTPContext{
			ctx: context.Background(),
			w:   v[0].(http.ResponseWriter),
			r:   v[1].(*http.Request),
		}
	}

	switch c := v[0].(type) {
	case echo.Context:
		return &EchoContext{ctx: c}
	}
	return nil
}

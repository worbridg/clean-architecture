package router

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

type EchoRouter struct {
	contexts map[string]reflect.Value
	e        *echo.Echo
	server   *http.Server
}

func NewEchoRouter(contexts Contexts) *EchoRouter {
	return &EchoRouter{
		contexts: contexts,
		e:        echo.New(),
		server:   &http.Server{Addr: ":3030"},
	}
}

func (e *EchoRouter) GET(path string, handler interface{}) {
	rf := reflect.ValueOf(handler)
	name := rf.Type().In(0).String()
	fn, ok := e.contexts[name]
	if !ok {
		panic("unknown context: " + name)
	}

	e.e.GET(path, func(c echo.Context) error {
		ctx := fn.Call([]reflect.Value{reflect.ValueOf(c)})
		rf.Call(ctx)
		return nil
	})
}

func (e *EchoRouter) Start() error {
	e.server.Handler = e.e
	return e.server.ListenAndServe()
}

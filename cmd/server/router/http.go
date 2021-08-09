package router

import (
	"net/http"
	"reflect"
)

type HTTPRouter struct {
	contexts map[string]reflect.Value
	mux      *http.ServeMux
	server   *http.Server
}

func NewHTTPRouter(contexts Contexts) *HTTPRouter {
	return &HTTPRouter{
		contexts: contexts,
		mux:      http.NewServeMux(),
		server:   &http.Server{Addr: ":3030"},
	}
}

func (r *HTTPRouter) GET(path string, handler interface{}) {
	rf := reflect.ValueOf(handler)
	fn := r.contexts[rf.Type().In(0).String()]

	r.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			ctx := fn.Call([]reflect.Value{
				reflect.ValueOf(w),
				reflect.ValueOf(r),
			})
			rf.Call(ctx)
		}
	})
}

func (r *HTTPRouter) Start() error {
	r.server.Handler = r.mux
	return r.server.ListenAndServe()
}

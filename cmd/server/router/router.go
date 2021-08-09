package router

type Router interface {
	GET(path string, handler interface{})
	Start() error
}

package user

type Handler interface {
	Get(ctx Context) (interface{}, error)
	Create(ctx Context) error
}

package user

type Usecase interface {
	GetUser(r *Request) (*Response, error)
	CreateUser(r *Request) error
}

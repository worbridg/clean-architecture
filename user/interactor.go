package user

type Interactor struct {
	repo Repository
}

func NewInteractor(repo Repository) *Interactor {
	return &Interactor{repo: repo}
}

func (i *Interactor) GetUser(r *Request) (*Response, error) {
	user, err := i.repo.GetByID(r.UserID)
	if err != nil {
		return nil, err
	}
	Logger.Printf("user: %v", user)

	return &Response{Name: user.Name()}, nil
}

func (i *Interactor) CreateUser(r *Request) error {
	Logger.Printf("creating new user: %v", r)
	return i.repo.Store(NewUser(0, r.UserName))
}

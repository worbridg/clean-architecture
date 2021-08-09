package repository

import (
	"errors"

	"github.com/worbridg/clean-architecture/user"
)

type InMemoryRepository struct {
	data map[user.UserID]*user.User
	i    user.UserID
}

var (
	ErrNotFoundUser = errors.New("not found user")
)

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		data: map[user.UserID]*user.User{},
	}
}

func (repo *InMemoryRepository) GetByID(id user.UserID) (*user.User, error) {
	user, ok := repo.data[id]
	if !ok {
		return nil, ErrNotFoundUser
	}
	return user, nil
}

func (repo *InMemoryRepository) Store(newUser *user.User) error {
	if newUser == nil {
		return ErrNotFoundUser
	}
	repo.i++
	repo.data[repo.i] = user.NewUser(repo.i, newUser.Name())
	Logger.Printf("stored in memory: %v", repo.data[repo.i])
	return nil
}

func (repo *InMemoryRepository) Delete(id user.UserID) error {
	delete(repo.data, id)
	return nil
}

package user

import "errors"

type UserName string

var (
	ErrEmptyUserName = errors.New("user name is empty")
)

func NewUserName(name string) (UserName, error) {
	if name == "" {
		return "", ErrEmptyUserName
	}
	return UserName(name), nil
}

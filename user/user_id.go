package user

import (
	"errors"
	"strconv"
)

type UserID int

var (
	ErrNegativeUserId = errors.New("user ID must be positive")
)

func NewUserID(id int) (UserID, error) {
	if id < 0 {
		return 0, ErrNegativeUserId
	}
	return UserID(id), nil
}

func (id UserID) String() string {
	return strconv.Itoa(int(id))
}

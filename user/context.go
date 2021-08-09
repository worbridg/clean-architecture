package user

import (
	"context"
)

type Context interface {
	Context() context.Context
	UserID() UserID
	UserName() UserName
	JSON(status int, v interface{})
}

package context

import (
	"context"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/worbridg/clean-architecture/user"
)

type EchoContext struct {
	ctx echo.Context
}

func (e *EchoContext) Context() context.Context {
	return context.Background()
}

func (e *EchoContext) UserID() user.UserID {
	n, err := strconv.Atoi(e.ctx.QueryParam("userId"))
	if err != nil {
		return 0
	}
	userId, err := user.NewUserID(n)
	if err != nil {
		return 0
	}

	return userId
}

func (e *EchoContext) UserName() user.UserName {
	userName, err := user.NewUserName(e.ctx.QueryParam("userName"))
	if err != nil {
		return ""
	}
	return userName
}

func (e *EchoContext) JSON(status int, v interface{}) {
	e.ctx.JSON(status, v)
}

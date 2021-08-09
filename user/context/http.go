package context

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/worbridg/clean-architecture/user"
)

type HTTPContext struct {
	ctx context.Context
	w   http.ResponseWriter
	r   *http.Request
}

func (c *HTTPContext) Context() context.Context {
	return c.ctx
}

func (c *HTTPContext) UserID() user.UserID {
	n, err := strconv.Atoi(c.r.FormValue("userId"))
	if err != nil {
		return 0
	}
	userId, err := user.NewUserID(n)
	if err != nil {
		return 0
	}
	return userId
}

func (c *HTTPContext) UserName() user.UserName {
	userName, err := user.NewUserName(c.r.FormValue("userName"))
	if err != nil {
		return ""
	}
	return userName
}

func (c *HTTPContext) JSON(status int, v interface{}) {
	json.NewEncoder(c.w).Encode(v)
}

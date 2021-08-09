package handler

import (
	"net/http"

	"github.com/worbridg/clean-architecture/user"
)

type HTTPHandler struct {
	usecase user.Usecase
}

func NewHTTPHandler(usecase user.Usecase) *HTTPHandler {
	return &HTTPHandler{
		usecase: usecase,
	}
}

func (h *HTTPHandler) Get(ctx user.Context) (interface{}, error) {
	req := &user.Request{
		UserID: ctx.UserID(),
	}
	res, err := h.usecase.GetUser(req)
	Logger.Printf("response: %v", res)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, struct {
			Err error `json:"error_message"`
		}{Err: err})
	} else {
		ctx.JSON(http.StatusOK, Model{User: User{Name: string(res.Name)}})
	}

	return res, nil
}

func (h *HTTPHandler) Create(ctx user.Context) error {
	req := &user.Request{UserName: ctx.UserName()}
	if err := h.usecase.CreateUser(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, struct {
			Err error `json:"error_message"`
		}{Err: err})
	} else {
		ctx.JSON(http.StatusOK, nil)
	}

	return nil
}

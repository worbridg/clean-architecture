package di

import (
	"os"

	"github.com/worbridg/clean-architecture/user"
	"github.com/worbridg/clean-architecture/user/context"
	"github.com/worbridg/clean-architecture/user/handler"
	"github.com/worbridg/clean-architecture/user/repository"
)

func UserHandler() user.Handler {
	user.Logger.SetOutput(os.Stdout)
	handler.Logger.SetOutput(os.Stdout)
	repository.Logger.SetOutput(os.Stdout)
	context.Logger.SetOutput(os.Stdout)
	return handler.NewHTTPHandler(user.NewInteractor(repository.NewInMemoryRepository()))
}

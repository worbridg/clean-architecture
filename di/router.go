package di

import (
	"reflect"

	"github.com/worbridg/clean-architecture/cmd/server/router"
	user "github.com/worbridg/clean-architecture/user/context"
)

func contexts() router.Contexts {
	return map[string]reflect.Value{
		"user.Context": reflect.ValueOf(user.NewContext),
	}
}

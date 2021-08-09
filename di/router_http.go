// +build http

package di

import (
	"github.com/worbridg/clean-architecture/cmd/server/router"
)

func Router() router.Router {
	return router.NewHTTPRouter(contexts())
}

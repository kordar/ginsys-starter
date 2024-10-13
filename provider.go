package ginsys_starter

import (
	"github.com/gin-gonic/gin"
	digstarter "github.com/kordar/dig-starter"
	"github.com/kordar/ginsys/middleware"
	"github.com/kordar/ginsys/resp"
	"github.com/kordar/gotrans"
	"go.uber.org/dig"
)

func serverProvider() {
	digstarter.ProvideE(func() (gin.HandlerFunc, error) {
		return middleware.CorsMiddleware(), nil
	}, dig.Group(ServerMiddlewareIocGroup.Value()))

	if I18nEnabled {
		digstarter.ProvideE(func() (gin.HandlerFunc, error) {
			return resp.TransLocaleMiddleware("Locale", "zh"), nil
		}, dig.Group(ServerMiddlewareIocGroup.Value()))

		digstarter.ProvideE(func() (gotrans.GoTranslation, error) {
			return gotrans.NewEnTranslation(), nil
		}, dig.Group(ServerTransIocGroup.Value()))

		digstarter.ProvideE(func() (gotrans.GoTranslation, error) {
			return gotrans.NewZhTranslation(), nil
		}, dig.Group(ServerTransIocGroup.Value()))
	}

}

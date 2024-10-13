package ginsys_starter

import (
	"github.com/gin-gonic/gin"
	digstarter "github.com/kordar/dig-starter"
	"github.com/kordar/ginsys"
	gocfgmodule "github.com/kordar/gocfg-load-module"
	"github.com/kordar/gotrans"
	"github.com/kordar/govalidator"
	"go.uber.org/dig"
)

type GinServerParam struct {
	dig.In
	Middleware    []gin.HandlerFunc          `group:"middleware"`
	Trans         []gotrans.GoTranslation    `group:"gotrans"`
	Validations   []govalidator.IValidation  `group:"govalidator"`
	RouterGroupFn []func(router *gin.Engine) `group:"router-group"`
}

func Server() {
	defer DestroyStarter()
	server := ginsys.NewGinServer()
	digstarter.InvokeE(func(p GinServerParam) {
		server.
			Middleware(p.Middleware...).
			OpenValidateAndTranslations(p.Trans...).
			AddValidate(p.Validations...)
		if p.RouterGroupFn != nil {
			for _, fn := range p.RouterGroupFn {
				fn(server.R())
			}
		}
	})
	server.Start()
}

func ServerStarter(starters []gocfgmodule.GoCfgModule) {
	InitStarter(starters)
	// 注入配置参数
	serverProvider()
	Server()
}

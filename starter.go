package ginsys_starter

import (
	"github.com/kordar/gocfg"
	gocfgmodule "github.com/kordar/gocfg-load-module"
)

func DestroyStarter() {
	gocfgmodule.Destroy()
}

func InitStarter(starters []gocfgmodule.GoCfgModule) {
	for _, starter := range starters {
		gocfgmodule.Register(starter)
	}
	// 解析组件
	gocfgmodule.ResolveAll(gocfg.AllSections())
}

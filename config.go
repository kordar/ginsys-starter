package ginsys_starter

import (
	"github.com/kordar/gocfg"
)

func InitDefaultConfig() {
	InitConfig("default", "./conf", "")
}

func InitDefaultI18nConfig() {
	InitConfig("default", "./conf", "language")
}

func InitConfig(group string, dir string, language string) {
	InitConfigWithParam(group, dir, language, []string{"ini", "yml", "yaml", "toml"})
}

func InitConfigWithParam(group string, dir string, language string, exts []string) {

	if language != "" {
		gocfg.InitConfigWithParentDir(language, exts...)
		I18nEnabled = true
	}

	gocfg.InitConfigWithDir(group, dir, exts...)
	ActiveGroup = group
}

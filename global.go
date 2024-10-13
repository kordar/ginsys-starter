package ginsys_starter

type IocGroup string

func (group IocGroup) Value() string {
	return string(group)
}

var (
	RouterIocGroup           IocGroup = "router-group"
	ServerMiddlewareIocGroup IocGroup = "middleware"
	ServerTransIocGroup      IocGroup = "gotrans"
	ServerValidateIocGroup   IocGroup = "govalidator"
)

var (
	ActiveGroup = "default"
	I18nEnabled bool
)

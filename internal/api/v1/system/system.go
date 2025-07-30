package system

import "github.com/rulessly/gin-base/internal/service/system"

type ApiGroup struct {
	BaseApi
}

var (
	systemUserService = system.NewSystemUserService()
)

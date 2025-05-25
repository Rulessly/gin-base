package v1

import "github.com/rulessly/gin-base/internal/api/v1/system"

var ApiGroup = new(apiGroup)

type apiGroup struct {
	SystemApiGroup system.ApiGroup
}

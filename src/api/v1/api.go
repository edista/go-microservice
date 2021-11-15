package v1

import "github.com/labstack/echo/v4"

type Api struct {
	version string
	e       *echo.Echo
}

func NewApi(e *echo.Echo) {
	api := Api{
		version: "v1",
		e:       e,
	}
	NewPostApi(api)
}

func (api *Api) CreateGroup(groupName string) *echo.Group {
	return api.e.Group("/api/" + api.version + "/" + groupName)
}

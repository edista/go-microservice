package v1

import (
	"github.com/edista/go-microservice/src/service"
	"github.com/labstack/echo/v4"
)

type Api struct {
	version string
	app     *echo.Echo
	service *service.Service
}

func NewApi(app *echo.Echo, service *service.Service) {
	api := Api{
		version: "v1",
		app:     app,
		service: service,
	}

	NewPostApi(api, service.PostService)
	NewUserApi(api)
}

func (api *Api) CreateGroup(groupName string) *echo.Group {
	return api.app.Group("/api/" + api.version + "/" + groupName)
}

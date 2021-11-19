package v1

import (
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func NewUserApi(api Api) {
	h := UserHandler{}
	g := api.CreateGroup("/users")
	g.GET("", h.listPost)
	g.GET("/:id", h.getPost)
	g.POST("", h.createPost)
	g.PUT("/:id", h.updatePost)
}

func (h *UserHandler) listPost(c echo.Context) error {
	return nil
}

func (h *UserHandler) getPost(c echo.Context) error {
	return nil
}

func (h *UserHandler) createPost(c echo.Context) error {
	return nil
}

func (h *UserHandler) updatePost(c echo.Context) error {
	return nil
}

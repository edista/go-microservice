package v1

import (
	"net/http"

	"github.com/edista/go-microservice/src/service/post"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	postService post.PostService
}

func NewPostApi(api Api, postService post.PostService) {
	h := PostHandler{
		postService: postService,
	}
	g := api.CreateGroup("/posts")
	g.GET("", h.listPost)
	g.GET("/:id", h.getPost)
	g.POST("", h.createPost)
	g.PUT("/:id", h.updatePost)
}

func (h *PostHandler) listPost(c echo.Context) error {
	u := h.postService.ListPost()
	return c.JSON(http.StatusOK, u)
}

func (h *PostHandler) getPost(c echo.Context) error {
	return nil
}

func (h *PostHandler) createPost(c echo.Context) error {
	return nil
}

func (h *PostHandler) updatePost(c echo.Context) error {
	return nil
}

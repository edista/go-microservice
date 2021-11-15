package v1

import (
	"net/http"

	"github.com/edista/go-microservice/src/service/post"
	"github.com/labstack/echo/v4"
)

type PostsApi struct {
}

func NewPostApi(api Api) {
	group := api.CreateGroup("posts")
	group.GET("", listPost)
	group.GET("/:id", getPost)
	group.POST("", createPost)
	group.PUT("/:id", updatePost)
}

func listPost(c echo.Context) error {
	u := post.ListPost()
	return c.JSON(http.StatusOK, u)
}

func getPost(c echo.Context) error {
	return nil
}

func createPost(c echo.Context) error {
	return nil
}

func updatePost(c echo.Context) error {
	return nil
}

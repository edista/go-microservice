package post

import "github.com/edista/go-microservice/src/model"

func ListPost() []model.Post {
	return []model.Post{
		{
			ID:        1,
			Title:     "Go Microservice By Edi",
			Content:   "Hello World",
			ViewCount: 1,
		},
	}
}

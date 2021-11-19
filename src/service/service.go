package service

import (
	"github.com/edista/go-microservice/engine"
	"github.com/edista/go-microservice/src/config"
	"github.com/edista/go-microservice/src/service/post"
)

type Service struct {
	PostService post.PostService
}

func NewService(engine *engine.Engine, config *config.Config) *Service {
	return &Service{
		PostService: *post.NewPostService(engine),
	}
}

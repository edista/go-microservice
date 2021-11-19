package post

import (
	"context"

	"github.com/edista/go-microservice/engine"
	"github.com/edista/go-microservice/src/model"
)

type PostService struct {
	engine *engine.Engine
}

func NewPostService(engine *engine.Engine) *PostService {
	return &PostService{engine: engine}
}

func (s *PostService) fetchAllPost(ctx context.Context) (posts []model.Post) {
	trx := s.engine.Database.WithTransaction(ctx)
	trx.Raw(`
		SELECT * FROM posts 
		WHERE deleted_at IS NULL
	`).Find(&posts)
	return posts
}

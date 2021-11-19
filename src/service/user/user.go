package post

import "github.com/edista/go-microservice/engine"

type UserService struct {
	Config   map[string]string
	Database int
}

func NewUserService(engine *engine.Engine) *UserService {
	return &UserService{}
}

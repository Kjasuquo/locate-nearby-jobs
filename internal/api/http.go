package api

import "github.com/kjasuquo/jobslocation/internal/repository"

type HTTPHandler struct {
	Repository repository.Repository
	RedisRepo  repository.RedisRepo
}

func NewHTTPHandler(repository repository.Repository, redisRepo repository.RedisRepo) *HTTPHandler {
	return &HTTPHandler{
		Repository: repository,
		RedisRepo:  redisRepo,
	}
}

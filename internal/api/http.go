package api

import "github.com/kjasuquo/jobslocation/internal/repository"

type HTTPHandler struct {
	Repository repository.Repository
}

func NewHTTPHandler(repository repository.Repository) *HTTPHandler {
	return &HTTPHandler{Repository: repository}
}

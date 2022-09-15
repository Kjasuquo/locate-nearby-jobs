package repository

import (
	"context"
	"github.com/kjasuquo/jobslocation/internal/model"
)

type Repository interface {
	SearchJobsByLocation(title string, long, lat float64) ([]model.Jobs, error)
	SearchJobsByTitle(title string) ([]model.Jobs, error)
}

type RedisRepo interface {
	Get(ctx context.Context, title string) ([]model.Jobs, error)
	Set(ctx context.Context, title string, job []model.Jobs) error
}

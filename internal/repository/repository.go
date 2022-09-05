package repository

import "github.com/kjasuquo/jobslocation/internal/model"

type Repository interface {
	SearchJobsByLocation(title string, long, lat float64) []model.Jobs
	SearchJobsByTitle(title string) ([]model.Jobs, error)
}

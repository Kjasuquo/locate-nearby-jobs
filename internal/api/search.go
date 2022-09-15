package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/kjasuquo/jobslocation/internal/helpers"
	"github.com/kjasuquo/jobslocation/internal/model"
	"net/http"
	"strconv"
)

//SearchByLocation uses you location to get available jobs with 5km radius
func (u *HTTPHandler) SearchByLocation(c *gin.Context) {

	longitude, err := strconv.ParseFloat(c.Query("long"), 64)
	if err != nil {
		helpers.Response(c, "Bad request", http.StatusBadRequest, nil, []string{"Bad Request"})
		return
	}

	latitude, err := strconv.ParseFloat(c.Query("lat"), 64)
	if err != nil {
		helpers.Response(c, "Bad request", http.StatusBadRequest, nil, []string{"Bad Request"})
		return
	}

	title := c.Query("title")

	//For testing: use longitude: 103.851 and latitude: 1.30156 and title: Developer

	jobs, err := u.Repository.SearchJobsByLocation(title, longitude, latitude)
	if err != nil {
		helpers.Response(c, "Internal Server Error", http.StatusInternalServerError, nil, []string{"Internal Server Error"})
		return
	}

	helpers.Response(c, "jobs successfully found", 200, jobs, nil)
}

//SearchByTitle gets all the jobs in the database regardless of location
func (u *HTTPHandler) SearchByTitle(c *gin.Context) {
	title := c.Query("title")

	var jobs []model.Jobs

	jobs, err := u.RedisRepo.Get(c, title)
	if err == redis.Nil {
		jobs, err = u.Repository.SearchJobsByTitle(title)
		if err != nil {
			helpers.Response(c, "An internal server error", 500, nil, []string{"internal server error"})
			return
		}

		err = u.RedisRepo.Set(c, title, jobs)
		if err != nil {
			helpers.Response(c, "An internal server error", 500, nil, []string{"error setting cache"})
			return
		}
	} else if err != nil {
		if err != nil {
			helpers.Response(c, "An internal server error", 500, nil, []string{"error getting data from cache"})
			return
		}
	}

	helpers.Response(c, "jobs successfully found", 200, jobs, nil)
}

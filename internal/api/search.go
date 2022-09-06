package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kjasuquo/jobslocation/internal/helpers"
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

	jobs := u.Repository.SearchJobsByLocation(title, longitude, latitude)

	helpers.Response(c, "jobs successfully found", 200, jobs, nil)
}

//SearchByTitle gets all the jobs in the database regardless of location
func (u *HTTPHandler) SearchByTitle(c *gin.Context) {
	title := c.Query("title")

	jobs, err := u.Repository.SearchJobsByTitle(title)
	if err != nil {
		helpers.Response(c, "An internal server error", 500, nil, []string{"internal server error"})
		return
	}

	helpers.Response(c, "jobs successfully found", 200, jobs, nil)
}

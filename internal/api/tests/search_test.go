package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/kjasuquo/jobslocation/cmd/server"
	"github.com/kjasuquo/jobslocation/internal/api"
	"github.com/kjasuquo/jobslocation/internal/model"
	"github.com/kjasuquo/jobslocation/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := mocks.NewMockRepository(ctrl)

	handler := &api.HTTPHandler{Repository: mockDB}
	router := server.SetupRouter(handler)

	job1 := model.Jobs{
		Title:     "Developer",
		Longitude: 103.851,
		Latitude:  1.30156,
	}

	job2 := model.Jobs{
		Title:     "Business",
		Longitude: 103.851,
		Latitude:  1.30156,
	}

	jobs := []model.Jobs{
		job1,
		job2,
	}

	bytes, _ := json.Marshal(jobs)

	job := []model.Jobs{
		job1,
	}

	var noJob []model.Jobs

	t.Run("SearchByLocation: testing successful request without title", func(t *testing.T) {
		mockDB.EXPECT().SearchJobsByLocation("", job1.Longitude, job1.Latitude).Return(jobs, nil)

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/?long=103.851&lat=1.30156&title=", strings.NewReader(string(bytes)))
		router.ServeHTTP(rw, req)
		fmt.Println(rw.Code)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "jobs successfully found")
	})

	t.Run("SearchByLocation: testing with no job in location", func(t *testing.T) {
		mockDB.EXPECT().SearchJobsByLocation("", 1.2345, 23.0000).Return(noJob, nil)

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/?long=1.2345&lat=23&title=", strings.NewReader(string(bytes)))
		router.ServeHTTP(rw, req)
		fmt.Println(rw.Code)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "jobs successfully found")
	})

	t.Run("SearchByLocation: testing successful request with title", func(t *testing.T) {
		mockDB.EXPECT().SearchJobsByLocation("Developer", job1.Longitude, job1.Latitude).Return(job, nil)

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/?long=103.851&lat=1.30156&title=Developer", strings.NewReader(string(bytes)))
		router.ServeHTTP(rw, req)
		fmt.Println(rw.Code)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "jobs successfully found")
	})

	t.Run("SearchByLocation: testing for error", func(t *testing.T) {
		mockDB.EXPECT().SearchJobsByLocation("Developer", job1.Longitude, job1.Latitude).Return(nil, errors.New("internal server error"))

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/?long=103.851&lat=1.30156&title=Developer", strings.NewReader(string(bytes)))
		router.ServeHTTP(rw, req)
		fmt.Println(rw.Code)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "Internal Server Error")
	})

	t.Run("SearchByTitle: testing successful request without title", func(t *testing.T) {
		mockDB.EXPECT().SearchJobsByTitle("").Return(jobs, nil)

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/title?title=", strings.NewReader(string(bytes)))
		router.ServeHTTP(rw, req)
		fmt.Println(rw.Code)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "jobs successfully found")
	})

	t.Run("SearchByTitle: testing successful request with title", func(t *testing.T) {
		mockDB.EXPECT().SearchJobsByTitle("Developer").Return(job, nil)

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/title?title=Developer", strings.NewReader(string(bytes)))
		router.ServeHTTP(rw, req)
		fmt.Println(rw.Code)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), "jobs successfully found")
	})

	t.Run("SearchByTitle: testing error", func(t *testing.T) {
		mockDB.EXPECT().SearchJobsByTitle("Developer").Return(nil, errors.New("internal server error"))

		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/title?title=Developer", strings.NewReader(string(bytes)))
		router.ServeHTTP(rw, req)
		fmt.Println(rw.Code)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "An internal server error")
	})
}

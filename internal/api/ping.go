package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kjasuquo/jobslocation/internal/helpers"
)

//PingHandler is for testing the connections
func (u *HTTPHandler) PingHandler(c *gin.Context) {
	// healthcheck
	helpers.Response(c, "pong", 200, nil, nil)
}

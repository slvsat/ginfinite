package ginfinite

import "github.com/gin-gonic/gin"

type Response struct {
	status int
	body   interface{}
}

type HandlerFunc func(*gin.Context) (*Response, error)

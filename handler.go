package ginfinite

import "github.com/gin-gonic/gin"

type HandlerFunc func(*gin.Context) (interface{}, error)

package ginfinite

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinfiniteServer struct {
	server *gin.Engine
}

func (g *GinfiniteServer) NewGinfinite() *GinfiniteServer {
	return &GinfiniteServer{
		server: gin.Default(),
	}
}

func (g *GinfiniteServer) Handle(method, relativePath string, handler HandlerFunc) {
	g.server.Handle(method, relativePath, ToPureGinJSON(handler))
}

func (g *GinfiniteServer) StartServer(port string) error {
	return http.ListenAndServe(":"+port, g.server)
}

func ToPureGinJSON(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		result, err := handler(c)
		if err != nil {
			ginifiniteErr, ok := err.(*GinfiniteError)
			if ok {
				c.JSON(ginifiniteErr.status, ginifiniteErr.message)
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
		return
	}
}

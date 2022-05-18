package ginServer

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// GinRoutingFn is callback function for setting up routers.
type GinRoutingFn func(router *gin.Engine)

// GinServerParams contains all dependencies of ginServer.
type Params struct {
	dig.In
	Routing GinRoutingFn
	Conf    *Config
	Router  *gin.Engine
}

// NewGinServer returns new instance of Server.
func NewGinServer(params Params) Server {
	return &ginServer{
		conf:    params.Conf,
		router:  params.Router,
		routing: params.Routing,
	}
}

type ginServer struct {
	routing GinRoutingFn
	conf    *Config
	router  *gin.Engine
}

func (g *ginServer) Open() error {
	if g.routing == nil {
		return ErrNilRoutingFn
	}
	g.routing(g.router)

	if err := g.router.Run(g.conf.ListenerAddr()); err != nil {
		return err
	}

	return nil
}

func (g *ginServer) Close() {
	// Blank function
}

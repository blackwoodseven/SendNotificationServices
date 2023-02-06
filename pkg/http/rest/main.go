package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	Engine() *gin.Engine
	Run() error
}

type service struct {
	conf   *Configuration
	engine *gin.Engine
}

type Configuration struct {
	Env  string
	Host string
	Port uint16
}

func Init(conf *Configuration) Repository {
	s := &service{
		conf:   conf,
		engine: gin.New(),
	}

	if conf.Env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	s.engine.SetTrustedProxies(nil)

	return s
}

// Engine returns the underlying server engine used by the HTTP API server.
func (s *service) Engine() *gin.Engine {
	return s.engine
}

// Run starts the HTTP API server. It returns an error if the server panics.
func (s *service) Run() error {
	return s.engine.Run(fmt.Sprintf("%s:%d", s.conf.Host, s.conf.Port))
}

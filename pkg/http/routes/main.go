package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/http/health"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/http/rest"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/http/rest/requestnotification"
)

type repository struct {
	Input
}

type Repository interface {
	Configure()
}

type Input struct {
	API                 rest.Repository
	Health              health.Repository
	RequestNotification requestnotification.Repository
}

func Init(input Input) Repository {
	return &repository{
		input,
	}
}

func (r *repository) Configure() {
	engine := r.API.Engine()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	api := engine.Group("/api")
	api.GET("/health", r.Health.GetHealth)
	api.POST("/requestnotification/sendrequest", r.RequestNotification.RequestNotification)
}

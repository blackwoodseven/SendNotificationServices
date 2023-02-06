package main

import (
	"fmt"
	"os"

	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/domain/requestnotification"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/http/health"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/http/rest"
	restapi "github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/http/rest/requestnotification"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/http/routes"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/logger"
	"github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/storage/db"
	notificationdb "github.com/rajivganesamoorthy-kantar/SendNotificationServices/pkg/storage/db/notificationrequest"
	"gorm.io/gorm"
)

type Storage struct {
	NotificationRequest notificationdb.Repository
}

func initStorage(conn *gorm.DB) Storage {
	logGroup := "Init Storage"
	notificationDb, err := notificationdb.InitDB(conn)
	if err != nil {
		logger.LogCriticalError(logGroup, fmt.Errorf("error: failed to init coffee date storage"))
		panic(err)
	}

	return Storage{
		NotificationRequest: notificationDb,
	}
}

type Domains struct {
	NotificationRequest requestnotification.Domain
}

func initDomains(s Storage) Domains {
	logGroup := "Init Domain"

	notificationDomain, err := requestnotification.Init(requestnotification.Input{
		NotificationRequest: s.NotificationRequest,
	})

	if err != nil {
		logger.LogCriticalError(logGroup, fmt.Errorf("error: failed to init notification domain"))
		panic(err)
	}
	return Domains{
		NotificationRequest: notificationDomain,
	}
}

type API struct {
	NotificationRequest restapi.Repository
	Health              health.Repository
	Rest                rest.Repository
}

func initAPIServices(d Domains) API {
	logGroup := "rest"
	// init rest service
	config := rest.Init(&rest.Configuration{
		Env:  getenv("ENVIRONMENT", "dev"),
		Host: getenv("HTTP_HOST", "0.0.0.0"),
		Port: 8080,
	})

	notification, err := restapi.Init(restapi.Input{
		NotificationRequest: d.NotificationRequest,
	})

	if err != nil {
		logger.LogCriticalError(logGroup, fmt.Errorf("error: failed to init notification request date rest"))
		panic(err)
	}

	return API{
		NotificationRequest: notification,
		Health:              health.Init(),
		Rest:                config,
	}
}

func getenv(key, fb string) string {
	v := os.Getenv(key)

	if v == "" {
		return fb
	}

	return v
}

func initWebServices(d Domains) {

}

func main() {
	logGroup := "main"
	logger.LogInfo(logGroup, "Starting up")

	//establishes db connection with the given paramters, using main.go in db package
	dbConnection := db.Init(db.Input{
		//for docker
		//Host: getenv("DB_HOST", "project_db"),
		//Port: getenv("DB_PORT", "3306"),
		//for local
		Host:     getenv("DB_HOST", "localhost"),
		Port:     getenv("DB_PORT", "5432"),
		User:     getenv("DB_USER", "testuser"),
		Password: getenv("DB_PASSWORD", "123456"),
		Database: getenv("DB_NAME", "notificationdb"),
		Env:      getenv("ENVIRONMENT", "dev"),
	})

	storage := initStorage(dbConnection)
	doms := initDomains(storage)
	go doms.NotificationRequest.Scheduler()
	api := initAPIServices(doms)

	routes := routes.Init(routes.Input{
		API:                 api.Rest,
		Health:              api.Health,
		RequestNotification: api.NotificationRequest,
	})

	routes.Configure()

	err := api.Rest.Run()
	if err != nil {
		logger.Log("main", logger.SeverityError, "server terminated with error", err, nil)
	}
}

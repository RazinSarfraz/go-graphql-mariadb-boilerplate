package service

import (
	"github.com/appointment-api/conf"
	mariadb "github.com/appointment-api/repository/mariaDB"
)

type Container struct {
	ConfigService    ConfigService
	Config           conf.Config
	JobLisingService JobLisingService
}

func NewServiceContainer() *Container {
	configService := NewConfigService()
	config := conf.GetConfig()
	store := mariadb.SharedStore()
	jobListingService := NewJobListingService(store)
	return &Container{
		ConfigService:    configService,
		Config:           *config,
		JobLisingService: jobListingService,
	}
}

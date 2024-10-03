package service

import "github.com/appointment-api/conf"

type ConfigService interface {
	GetConfig() *conf.Config
}
type configService struct {
}

func NewConfigService() ConfigService {
	return &configService{}
}

func (c *configService) GetConfig() *conf.Config {

	return conf.GetConfig()

}

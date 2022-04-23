package config

import "github.com/penk110/interview_ext_go/z_ioc_gin/service"

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) NewOrder() *service.Order {
	return service.NewOrder()
}

func (config *Config) NewDB() *service.DB {
	return service.NewDB()
}

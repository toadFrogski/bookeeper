package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

type ApplicationConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string `default:"0.0.0.0"`
	Port int    `default:"8080"`
}

type DatabaseConfig struct {
	Host     string `default:"0.0.0.0"`
	Port     int    `default:"3306"`
	Username string
	Password string
	Database string
}

func (sc *ServerConfig) GetListenAddr() string {
	return fmt.Sprintf("%s:%d", sc.Host, sc.Port)
}

func (dc *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dc.Username, dc.Password, dc.Host, dc.Port, dc.Database,
	)
}

func LoadConfig(path string) (*ApplicationConfig, error) {
	var applicationConfig ApplicationConfig
	if err := configor.Load(&applicationConfig, path); err != nil {
		return nil, err
	}

	return &applicationConfig, nil
}

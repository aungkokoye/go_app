package services

import (
	"io/ioutil"

	"github.com/aungkokoye/go_app/databases"
	"gopkg.in/yaml.v3"
)

type Config struct {
	MysqlDB  databases.ConfigMysql `yaml:"mysql"`
	LogLevel string                `yaml:"loglevel"`
}

func NewConif(fileNmae string) (Config, error) {

	var config Config

	// Read the config file
	data, err := ioutil.ReadFile(fileNmae)

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)

	return config, err

}

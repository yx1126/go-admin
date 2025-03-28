package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Mysql struct {
	Ip           string `yaml:"ip" json:"ip"`
	Port         int    `yaml:"port" json:"port"`
	User         string `yaml:"user" json:"user"`
	Password     string `yaml:"password" json:"password"`
	Database     string `yaml:"database" json:"database"`
	MaxIdleConns int    `yaml:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns" json:"maxOpenConns"`
}

type Server struct {
	Port int `yaml:"port" json:"port"`
}

type ConfigContext struct {
	Mysql  Mysql  `yaml:"mysql" json:"mysql"`
	Server Server `yaml:"server" json:"server"`
}

var Config *ConfigContext

func init() {
	conf, err := os.ReadFile("./conf.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(conf, &Config)
	if err != nil {
		panic(err)
	}
}

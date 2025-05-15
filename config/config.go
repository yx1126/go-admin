package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type MysqlConfig struct {
	Ip           string `yaml:"ip" json:"ip"`
	Port         int    `yaml:"port" json:"port"`
	User         string `yaml:"user" json:"user"`
	Password     string `yaml:"password" json:"password"`
	Database     string `yaml:"database" json:"database"`
	MaxIdleConns int    `yaml:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns" json:"maxOpenConns"`
}

type ServerConfig struct {
	Port int `yaml:"port" json:"port"`
}

type UserConfig struct {
	Password string `yaml:"password" json:"password"`
}

type RedisConfig struct {
	Ip       string `yaml:"ip" json:"ip"`
	Port     int    `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	Database int    `yaml:"database" json:"database"`
}

type ConfigContext struct {
	Mysql  MysqlConfig  `yaml:"mysql" json:"mysql"`
	Server ServerConfig `yaml:"server" json:"server"`
	User   UserConfig   `yaml:"user" json:"user"`
	Redis  RedisConfig  `yaml:"redis" json:"redis"`
}

var Mysql *MysqlConfig
var Server *ServerConfig
var User *UserConfig
var Redis *RedisConfig

func init() {
	var Config *ConfigContext
	conf, err := os.ReadFile("./conf.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(conf, &Config)
	if err != nil {
		panic(err)
	}
	Mysql = &Config.Mysql
	Server = &Config.Server
	User = &Config.User
	Redis = &Config.Redis
}

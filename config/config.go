package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// system

type SystemConfig struct {
	Name string `yaml:"name" json:"name"`
}

// mysql
type MysqlConfig struct {
	Ip           string `yaml:"ip" json:"ip"`
	Port         int    `yaml:"port" json:"port"`
	User         string `yaml:"user" json:"user"`
	Password     string `yaml:"password" json:"password"`
	Database     string `yaml:"database" json:"database"`
	MaxIdleConns int    `yaml:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns" json:"maxOpenConns"`
}

// redis
type RedisConfig struct {
	Ip       string `yaml:"ip" json:"ip"`
	Port     int    `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	Database int    `yaml:"database" json:"database"`
}

// server
type ServerConfig struct {
	Port int    `yaml:"port" json:"port"`
	Mode string `yaml:"mode" json:"mode"`
}

// user
type UserConfig struct {
	Password      string `yaml:"password" json:"password"`
	MaxRetryCount int    `yaml:"maxRetryCount" json:"maxRetryCount"`
	LockTime      int    `yaml:"lockTime" json:"lockTime"`
}

// token
type TokenConfig struct {
	Header     string `yaml:"header" json:"header"`
	Secret     string `yaml:"secret" json:"secret"`
	ExpireTime int    `yaml:"expireTime" json:"expireTime"`
}

// minio
type MinioConfig struct {
	Ip     string `yaml:"ip" json:"ip"`
	Port   int    `yaml:"port" json:"port"`
	Access string `yaml:"access" json:"access"`
	Secret string `yaml:"secret" json:"secret"`
}

type ConfigContext struct {
	System SystemConfig `yaml:"system" json:"system"`
	Mysql  MysqlConfig  `yaml:"mysql" json:"mysql"`
	Server ServerConfig `yaml:"server" json:"server"`
	User   UserConfig   `yaml:"user" json:"user"`
	Redis  RedisConfig  `yaml:"redis" json:"redis"`
	Token  TokenConfig  `yaml:"token" json:"token"`
	Minio  MinioConfig  `yaml:"minio" json:"minio"`
}

var System *SystemConfig
var Mysql *MysqlConfig
var Server *ServerConfig
var User *UserConfig
var Redis *RedisConfig
var Token *TokenConfig
var Minio *MinioConfig

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
	System = &Config.System
	Mysql = &Config.Mysql
	Server = &Config.Server
	User = &Config.User
	Redis = &Config.Redis
	Token = &Config.Token
	Minio = &Config.Minio
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// IV i like, i love
type IV struct {
	JwtSecret string `yaml:"jwt_secret"`
}

// Server server配置
type Server struct {
	RunMode      string `yaml:"run_mode"`
	HTTPPort     int    `yaml:"http_port"`
	ReadTimeout  int    `yaml:"read_timeout"`
	WriteTimeout int    `yaml:"write_timeout"`
}

// Database 数据库
type Database struct {
	Type        string `yaml:"type"`
	User        string `yaml:"user"`
	Password    string `yaml:"string"`
	Host        string `yaml:"host"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"table_prefix"`
}

// Redis redis数据库
type Redis struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"password"`
	MaxIdle     int    `yaml:"max_idle"`
	MaxActive   int    `yaml:"max_active"`
	IdleTimeout int    `yaml:"idle_timeout"`
}

// Env 环境变量
type Env struct {
	IV       IV       `yaml:"iv"`
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
}

func main() {
	fileData, err := ioutil.ReadFile("conf/config.yaml")
	if err != nil {
		log.Fatal("read conf/config.yaml fail:", err)
	}

	env := Env{}
	err = yaml.Unmarshal(fileData, &env)
	fmt.Println(env.IV)
	fmt.Println(env.Server)
	fmt.Println(env.Database)
	fmt.Println(env.Redis)
}

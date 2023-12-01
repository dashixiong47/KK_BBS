package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Settings struct {
	Application struct {
		Mode         string `yaml:"mode"`
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		SigningKey   string `yaml:"signingKey"`
		TokenTimeout int    `yaml:"tokenTimeout"`
		IdConfusion  int    `yaml:"idConfusion"`
		I18          string `yaml:"i18"`
		RedeemKey    string `yaml:"redeemKey"`
		FileHost     string `yaml:"fileHost"`
	} `yaml:"application"`
	Postgresql struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"postgresql"`
	Redis struct {
		Host               string `yaml:"host"`
		Port               int    `yaml:"port"`
		Password           string `yaml:"password"`
		Database           int    `yaml:"database"`
		PooSize            int    `yaml:"pooSize"`
		MinIdleConns       int    `yaml:"minIdleConns"`
		MaxConnAge         int    `yaml:"maxConnAge"`
		PoolTimeout        int    `yaml:"poolTimeout"`
		IdleTimeout        int    `yaml:"idleTimeout"`
		IdleCheckFrequency int    `yaml:"idleCheckFrequency"`
	} `yaml:"redis"`
	Es struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Index    string `yaml:"index"`
	} `yaml:"es"`

	Smtp struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"smtp"`
}

var SettingsConfig = new(Settings)
var SystemConfig = make(map[string]string)

func init() {
	// 读取 YAML 文件
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}
	// 创建 Settings 结构体的实例
	var settings Settings
	// 解析 YAML 数据
	err = yaml.Unmarshal(data, &settings)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}
	SettingsConfig = &settings
}

package config

import (
	"io/ioutil"
	"encoding/json"
)

type Host struct {
	Host     string
	Username string
	Password string
	Port     int
}

type Config struct {
	Hosts      []*Host
	Globalpwd  string
	Globalname string
	Globalport int
}

func NewConfig() *Config {
	c := &Config{
		Globalport: 22,
	}
	return c;
}

func (config *Config) ParseFile(configFile string) {
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(content, config)
}

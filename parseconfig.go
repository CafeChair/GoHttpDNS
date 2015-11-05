package main

import (
	"io/ioutil"
)

type RedisConfig struct {
	Addr string
	Port string
}

type WebServer struct {
	Addr string
	Port string
}

type Config struct {
	Redis      RedisConfig
	Dnsservers []string
	Webserver  WebServer
}

func (rc *RedisConfig) isValid() bool {
	return len(rc.Addr) > 0 && rc.Port == "6379"
}

func ParseConfig(filepath string) error {
	if config, err := ioutil.ReadFile(filepath); err != nil {
		return err
	}
	return nil
}

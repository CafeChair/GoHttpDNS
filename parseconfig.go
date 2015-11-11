package main

type RedisConfig struct {
	Addr string
	Port string
}

type WebServer struct {
	Addr string
	Port string
}

type Config struct {
	Redis     RedisConfig
	Dnsserver string
	Webserver WebServer
}

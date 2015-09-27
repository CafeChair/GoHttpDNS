package main

type RedisConfig struct {
	Addr string
	Port string
}

type DnsServer struct {
	Addr []string
	Port string
}

type WebServer struct {
	Addr string
	Port string
}


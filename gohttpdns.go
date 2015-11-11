package main

import (
	"github.com/Unknwon/goconfig"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"os"
)

func main() {
	configfile := os.Args[1]
	config := goconfig.LoadConfigFile(configfile)
	redisaddr, _ := config.GetValue("redis", "address")
	redisport, _ := config.GetValue("redis", "port")
	httpaddr, _ := config.GetValue("webserver", "address")
	httpport, _ := config.GetValue("webserver", "port")
	dnsaddr, _ := config.GetValue("dnsserver", "address")

	listenaddr := httpaddr + httpport
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/resolve/#domain", resolveHandler),
		rest.Get("/health", healthHandler),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(listenaddr, api.MakeHandler()))
}

package main

import (
	"github.com/ant0ine/go-json-rest/rest"
)

type Resp struct {
	Succes int
	Domain string
	IpAddr string
}

const (
	SUCC = 0
	FAIL = 1
)

func healthHandler(w rest.ResponseWriter, rep *rest.Request) {
	resp := "I am OK"
	w.WriteJson(&resp)
}

func resolveHandler(w rest.ResponseWriter, req *rest.Request) {
	domain := req.PathParam("domain")
	ipaddr, err := resolveFromRedis(domain)
	if err == nil {
		resp := Resp{
			Succes: SUCC,
			Domain: domain,
			IpAddr: ipaddr,
		}
		w.WriteJson(&resp)
	}
	ipaddr, errs := resolveFromDNS(domain)
	if errs != nil {
		resp := Resp{
			Succes: FAIL,
			Domain: domain,
			IpAddr: "resolve error, maybe this domain has not resolved",
		}
		w.WriteJson(&resp)
	} else {
		resp := Resp{
			Succes: SUCC,
			Domain: domain,
			IpAddr: ipaddr,
		}
		cacheRespToRedis(domain, ipaddr)
		w.WriteJson(&resp)
	}

}

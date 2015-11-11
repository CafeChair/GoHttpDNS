package main

import (
	"github.com/miekg/dns"
)

/*
第一次获取解析结果,从redis中获取，如果redis中没有,则从DNS Server解析
*/
func resolveFromDns(domain string) ([]string, error) {
	answer := make([]string, 0)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	c := new(dns.Client)
	in, _, err := c.Exchange(m, dnsserver+":53")
	if err != nil {
		return answer, err
	}
	for _, ain := range in.Answer {
		if a, ok := ain.(*dns.A); ok {
			answer = append(answer, a.A.String())
		}
	}
	return answer, nil
}

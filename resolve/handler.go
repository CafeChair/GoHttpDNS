package resolve
import (
	"fmt"
	"github.com/hoisie/web"
)

type Response struct {
	Code bool
	RealIP string
	DomainName string
	Msg string	
}

func Ping(ctx *web.Context) string {
	ret := "ok"
	return ret
}

func Resolve(ctx *web.Context) string {
	domain := ctx.Params["domain"]
	realip,domainname,err := GetResultFromCache(domain)
	if err == nil {
		resp := Response{
			Code: true,
			RealIP:realip,
			DomainName:domainname,
		}
		return resp.jsonString()
	}
	realip,domainname,err := DNSResolve(domain)
	if err != nil {
		resp:=Response{
			Code:false,
			Msg: fmt.Printf("%s", err)
		}
		return resp.jsonString()
	}
	resp := Response{
		Code:true,
		RealIP: *realip,
		DomainName: *domainname,
	}
	CacheToRedis(domain,*domainname,*realip)
	return resp.jsonString()
}

func GetResultFromCache(domain string) (realip, domainname string, error) {
	conn := redispool.Get()
	if conn == nil {
		return
	}
	defer conn.Close()
	
}

func DNSResolve(domain string) (*string, *string, error) {
	
}
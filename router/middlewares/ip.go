package middlewares

import (
	"context"
	"net"
	"net/http"
	"strings"
)

const (
	IpParamNameDefault = "ip"

	headerNameRealIp       = "X-REAL-IP"
	headerNameForwardedFor = "X-FORWARDED-FOR"
)

type IpOption func(ip *IP) *IP

type IP struct {
	name string
}

func NewIP(options ...IpOption) Middleware {
	middleware := &IP{}

	options = append([]IpOption{IpWithName(IpParamNameDefault)}, options...)

	for _, option := range options {
		middleware = option(middleware)
	}

	return middleware
}

func (middleware *IP) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		address := request.Header.Get(headerNameRealIp)
		if strings.TrimSpace(address) == "" {
			address = request.Header.Get(headerNameForwardedFor)
		}
		if strings.TrimSpace(address) == "" {
			address = request.RemoteAddr
		}

		next.ServeHTTP(
			writer,
			request.WithContext(
				context.WithValue(request.Context(), middleware.name, net.ParseIP(strings.Split(address, ":")[0])),
			))
	})
}

func IpWithName(name string) IpOption {
	return func(ip *IP) *IP {
		ip.name = name

		return ip
	}
}

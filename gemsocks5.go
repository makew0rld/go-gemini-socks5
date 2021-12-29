package gemsocks5

import (
	"net"

	"golang.org/x/net/proxy"
)

// ProxyFunc takes the SOCKS5 proxy address and authentication, and returns a function
// that can be used in the go-gemini Client config. Authentication is optional, and auth
// can be set to nil to disable it.
func ProxyFunc(proxyAddr string, auth *proxy.Auth) func(*net.Dialer, string) (net.Conn, error) {
	return func(dialer *net.Dialer, address string) (net.Conn, error) {
		d, err := proxy.SOCKS5("tcp", proxyAddr, auth, dialer)
		if err != nil {
			return nil, err
		}
		return d.Dial("tcp", address)
	}
}

package utils

// TODO: use validator: https://github.com/asaskevich/govalidator

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Proxy struct {
	Scheme   string // `json:"Scheme"`
	Username string // `json:"Username"`
	Password string // `json:"Password"`
	IPAddr   string // `json:"IPAddr"`
	Port     int // `json:"Port"`
	RDNS     bool // `json:"RDNS"`
}

// TODO: are there any other more elegant unified way ???
var (
	ErrBadFormat   error = errors.New("bad url format")
	ErrBadScheme   error = errors.New("bad url scheme")
	ErrBadUsername error = errors.New("bad url username")
	ErrBadPassword error = errors.New("bad url password")
	ErrBadIPAddr   error = errors.New("bad url ipaddr")
	ErrBadPort     error = errors.New("bad url port")
)

func parsePort(s string) (int, error) {
	var port int
	p, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		// fmt.Printf("%v\n", err)
		return port, ErrBadPort
	}
	return int(p), nil
}

// Parse Proxy struct from url string
func FromUrl(url string, rdns bool) (Proxy, error) {

	pattern_short := regexp.MustCompile(`(http|https|socks5|socks5h):\/\/([\d]+\.[\d]+\.[\d]+\.[\d]+):([\d]+)`)

	pattern_long := regexp.MustCompile(`(http|https|socks5|socks5h):\/\/([\w\_]+):([\w\_]+)@([\d]+\.[\d]+\.[\d]+\.[\d]+):([\d]{2,5})`)

	var proxy = Proxy{
		Scheme:   "",
		Username: "",
		Password: "",
		IPAddr:   "",
		Port:     0,
		RDNS:     rdns,
	}
	var err error

	if strings.Contains(url, "@") {
		match_long := pattern_long.FindStringSubmatch(url)
		if match_long == nil {
			return proxy, ErrBadFormat
		}
		// fmt.Printf("match_long: %#v\n", match_long)

		proxy.Scheme = match_long[1]
		proxy.Username = match_long[2]
		proxy.Password = match_long[3]

		proxy.IPAddr = match_long[4]
		proxy.Port, err = parsePort(match_long[5])
		if err != nil {
			return proxy, ErrBadPort
		}

	} else {
		match_short := pattern_short.FindStringSubmatch(url)
		if match_short == nil {
			return proxy, ErrBadFormat
		}
		// fmt.Printf("match_short: %#v\n", match_short)

		proxy.Scheme = match_short[1]
		proxy.Username = ""
		proxy.Password = ""
		proxy.IPAddr = match_short[2]
		proxy.Port, err = parsePort(match_short[3])
		if err != nil {
			return proxy, ErrBadPort
		}
	}

	// Ref:
	// - https://github.com/lanlin/notes/issues/109
	// - https://superuser.com/questions/1762341/does-chromium-not-support-socks5h-with-the-h-in-the-end
	//
	// Validate scheme
	allowed_schemes := []string{"http", "https", "socks5", "socks5h"}
	if !slices.Contains(allowed_schemes, proxy.Scheme) {
		return proxy, ErrBadScheme
	}

	// Validate username
	pattern_username := regexp.MustCompile(`^([a-zA-Z0-9_]+)$`)
	if proxy.Username!="" && !pattern_username.MatchString(proxy.Username) {
		return proxy, ErrBadUsername
	}

	// Validate password
	pattern_password := regexp.MustCompile(`^([a-zA-Z0-9!~@#$%^&*()_+]+)$`)
	if proxy.Password!="" && !pattern_password.MatchString(proxy.Password) {
		return proxy, ErrBadUsername
	}

	// Ref: https://stackoverflow.com/questions/5284147/validating-ipv4-addresses-with-regexp
	//
	// Validate ipaddr
	pattern_ipaddr_strict := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]|[0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9][0-9]|[0-9])$`)
	if !pattern_ipaddr_strict.MatchString(proxy.IPAddr) {
		return proxy, ErrBadIPAddr
	}

	// Validate prot number
	if proxy.Port > 65535 {
		return proxy, ErrBadPort
	}
	if proxy.Port < 1023 {
		if proxy.Port != 80 && proxy.Port != 443 {
			return proxy, ErrBadPort
		}
	}

	// All seems good so far
	return proxy, nil

}

// Generate url string from Proxy struct
func (p *Proxy) ToUrl() string {
	if len(p.Username) > 0 && len(p.Password) > 0 {
		return fmt.Sprintf("%s://%s:%s@%s:%d", p.Scheme, p.Username, p.Password, p.IPAddr, p.Port)
	}
	return fmt.Sprintf("%s://@%s:%d", p.Scheme, p.IPAddr, p.Port)
}

// Check if proxy is HTTP
func (p *Proxy) IsHTTP() bool {
	return p.Scheme == "http"
}

// Check if proxy is HTTPS
func (p *Proxy) IsHTTPS() bool {
	return p.Scheme == "https"
}

// Check if proxy is Socks5
func (p *Proxy) IsSocks5() bool {
	return p.Scheme == "socks5" || p.Scheme == "socks5h"
}

// Check if proxy is HTTP or HTTPS
func (p *Proxy) IsHTTP_OR_HTTPS() bool {
	return p.Scheme == "http" || p.Scheme == "https"
}

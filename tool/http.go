package tool

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

var (
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Duration(60) * time.Second,
	}
)

func Get(url string) (io.ReadCloser, error) {
	c := client
	resp, err := c.Get(url)
	time.Sleep(200 * time.Millisecond)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}

// Sets the client using the given proxy address
// Proxy addresses can take the following forms:
//   - http://address:port
//   - socks5://address:port
//   - socks5://username:password@address:port
func SetProxyClient(proxyAddress string) error {
	var proxyClient *http.Client
	proxyParsed, err := url.Parse(proxyAddress)
	if err != nil {
		return fmt.Errorf("couldn't parse proxy address: %v", proxyAddress)
	}

	switch proxyParsed.Scheme {
	case "socks5":
		dialer, err := proxy.FromURL(proxyParsed, proxy.Direct)
		if err != nil {
			return fmt.Errorf("couldn't create dailer from proxy address: %v", proxyAddress)
		}

		proxyClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial:  dialer.Dial,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: time.Duration(60) * time.Second,
		}
	case "http", "https":
		proxyClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyParsed),
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: time.Duration(60) * time.Second,
		}
	default:
		return fmt.Errorf("proxy scheme %v is unsupported", proxyParsed.Scheme)
	}

	client = proxyClient
	return nil
}

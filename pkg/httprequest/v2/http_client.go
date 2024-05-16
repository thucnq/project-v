package httprequest

import (
	"net"
	"net/http"
	"time"
)

type HttpClient struct {
	*http.Client
}

func NewClient(clients ...*http.Client) *HttpClient {
	var client *http.Client
	if len(clients) == 0 {
		transport := &http.Transport{
			DialContext: (&net.Dialer{
				// limits the time spent establishing a TCP connection
				Timeout: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			// MaxIdleConnsPerHost:   2000,
			// MaxIdleConns:          2000,
			// MaxConnsPerHost:       2000,
		}
		client = &http.Client{
			Transport: transport,
		}
	} else {
		client = clients[0]
	}
	return &HttpClient{
		client,
	}
}

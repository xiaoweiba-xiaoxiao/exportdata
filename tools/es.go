package tools

import(
	"net"
	"time"
	"net/http"
	"crypto/tls"
	es "github.com/elastic/go-elasticsearch/v6"
)


func NewEsclient(hosts []string)(*es.Client, error){
	escf := es.Config{
		Addresses: hosts,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   100,
			ResponseHeaderTimeout: 10 * time.Second,
			DialContext:           (&net.Dialer{Timeout: 10 * time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	}
	return es.NewClient(escf)
}

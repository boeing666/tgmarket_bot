package parser

import (
	"net/http"
	"net/http/cookiejar"
	"sync"
	"time"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

var (
	tlsclient  cycletls.CycleTLS
	httpclient *http.Client
	cookies    *cookiejar.Jar
	once       sync.Once
)

func initRequests() {
	once.Do(func() {
		tlsclient = cycletls.Init()
		cookies, _ = cookiejar.New(nil)
		httpclient = &http.Client{
			Jar:     cookies,
			Timeout: 10 * time.Second,
		}
	})
}

func tlsClient() cycletls.CycleTLS {
	initRequests()
	return tlsclient
}

func httpClient() *http.Client {
	initRequests()
	return httpclient
}

func tlsRequest(url string, headers map[string]string, body string, method string) (cycletls.Response, error) {
	return tlsClient().Do(url, cycletls.Options{
		DisableRedirect: true,
		Headers:         headers,
		Body:            body,
		Ja3:             "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,17513-27-5-65281-11-16-45-13-51-10-65037-35-43-23-18-0-41,29-23-24,0",
		UserAgent:       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
	}, method)
}

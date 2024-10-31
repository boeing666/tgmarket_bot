package parser

import (
	"sync"

	"github.com/Danny-Dasilva/CycleTLS/cycletls"
)

var (
	net  cycletls.CycleTLS
	once sync.Once
)

func netClient() cycletls.CycleTLS {
	once.Do(func() {
		net = cycletls.Init()
	})
	return net
}

func request(url string, headers map[string]string, body string, method string) (cycletls.Response, error) {
	return netClient().Do(url, cycletls.Options{
		Headers:   headers,
		Body:      body,
		Ja3:       "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,17513-27-5-65281-11-16-45-13-51-10-65037-35-43-23-18-0-41,29-23-24,0",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
	}, method)
}

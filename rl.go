package main

import (
	"net/http"
	"sync"
)

var dafaultClientKey = "default"

type RateLimiter struct {
	httpCli *http.Client
	// мапа с рейтами по каждому юзеру/ключу
	clientsRates sync.Map
	// рейт лимит
	limit int
	// буффер для переиспользования рейтов
	seriesPool *seriesPool
}

func (rl *RateLimiter) DoWithKey(req *http.Request, clientKey any) (*http.Response, error) {
	if rl.limit <= 0 {
		return rl.do(req)
	}

	//start := time.Now().Add(-time.Second)
	return nil, nil
}

func (rl *RateLimiter) do(req *http.Request) (*http.Response, error) {
	return rl.httpCli.Do(req)
}

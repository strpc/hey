package main

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func generateRequest(requestParam request) *http.Request {
	req, err := http.NewRequest(
		strings.ToUpper(requestParam.Method),
		requestParam.Url,
		nil,
	)
	if err != nil {
		usageAndExit("invalid request param")
	}

	header := make(http.Header)
	header.Set("Content-Type", *contentType)
	if len(requestParam.Headers) > 0 {
		for key, value := range requestParam.Headers {
			header.Set(key, value)
		}
	}
	req.Header = header

	if len(requestParam.Cookies) > 0 {
		for key, value := range requestParam.Cookies {
			req.AddCookie(&http.Cookie{
				Name:  key,
				Value: value,
			})
		}
	}

	if requestParam.Username != "" || requestParam.Password != "" {
		req.SetBasicAuth(requestParam.Username, requestParam.Password)
	}

	if requestParam.Host != "" {
		req.Host = requestParam.Host
	}

	if requestParam.Body != nil {
		data, err := json.Marshal(requestParam.Body)
		if err != nil {
			usageAndExit(err.Error())
		}
		req.Body = io.NopCloser(bytes.NewReader(data))
	}
	return req
}

func requestFactory(config config) func() *http.Request {
	rand.Seed(time.Now().UnixNano())
	return func() *http.Request {
		index := rand.Intn(len(config.Requests))
		randomRequestParam := config.Requests[index]
		return generateRequest(randomRequestParam)
	}
}

package isevengo

import (
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
}

type Response struct {
	Iseven bool
	Ad     string
	Error  string
}

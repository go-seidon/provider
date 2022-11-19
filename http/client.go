package http

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
)

type Client interface {
	Get(ctx context.Context, p RequestParam) (*RequestResult, error)
	Post(ctx context.Context, p RequestParam) (*RequestResult, error)
	Put(ctx context.Context, p RequestParam) (*RequestResult, error)
	Patch(ctx context.Context, p RequestParam) (*RequestResult, error)
	Delete(ctx context.Context, p RequestParam) (*RequestResult, error)
}

type RequestParam struct {
	Url    string
	Body   io.Reader
	Header map[string][]string
}

type RequestResult struct {
	StatusCode    int
	Body          io.ReadCloser
	ContentLength int64
	Header        map[string][]string
}

type httpClient struct {
	client *http.Client
}

func (h *httpClient) Get(ctx context.Context, p RequestParam) (*RequestResult, error) {
	return h.request(ctx, http.MethodGet, p)
}

func (h *httpClient) Post(ctx context.Context, p RequestParam) (*RequestResult, error) {
	return h.request(ctx, http.MethodPost, p)
}

func (h *httpClient) Put(ctx context.Context, p RequestParam) (*RequestResult, error) {
	return h.request(ctx, http.MethodPut, p)
}

func (h *httpClient) Patch(ctx context.Context, p RequestParam) (*RequestResult, error) {
	return h.request(ctx, http.MethodPatch, p)
}

func (h *httpClient) Delete(ctx context.Context, p RequestParam) (*RequestResult, error) {
	return h.request(ctx, http.MethodDelete, p)
}

func (h *httpClient) request(ctx context.Context, method string, p RequestParam) (*RequestResult, error) {
	req, err := http.NewRequestWithContext(ctx, method, p.Url, p.Body)
	if err != nil {
		return nil, err
	}

	if len(p.Header) == 0 {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header = p.Header
	}

	doRes, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	res := &RequestResult{
		StatusCode:    doRes.StatusCode,
		Body:          doRes.Body,
		ContentLength: doRes.ContentLength,
		Header:        doRes.Header,
	}
	return res, nil
}

func NewClient(opts ...ClientOption) *httpClient {
	p := ClientParam{}
	for _, opt := range opts {
		opt(&p)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	if p.ShouldCheckSSL {
		tr.TLSClientConfig.InsecureSkipVerify = false
	}

	return &httpClient{
		client: &http.Client{
			Transport: tr,
		},
	}
}

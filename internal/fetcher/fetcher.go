package fetcher

import (
	"api-spammer/internal/config"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Destination struct {
	Id       int
	Endpoint config.Endpoint
}

type FetchResult struct {
	Id              int
	StatusCode      int
	ResponseHeaders http.Header
	Body            string
	ReplyTime       int
}

var ErrorResult = FetchResult{StatusCode: 0, Body: "", ReplyTime: 0}

func addHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Add(k, v)
	}
}

func createGetReq(ctx context.Context, d *Destination) *http.Request {
	query := url.Values{}
	for k, v := range d.Endpoint.Data {
		query.Add(k, v.(string))
	}
	baseUrl := d.Endpoint.Url
	fullUrl := baseUrl + "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, d.Endpoint.Method, fullUrl, nil)
	if err != nil {
		panic(err)
	}

	addHeaders(req, d.Endpoint.Headers)
	return req
}
func createPostReq(ctx context.Context, d *Destination) *http.Request {
	reqBytes, err := json.Marshal(d.Endpoint.Data)
	if err != nil {
		panic(err)
	}
	reqBody := bytes.NewBuffer(reqBytes)

	req, err := http.NewRequestWithContext(ctx, d.Endpoint.Method, d.Endpoint.Url, reqBody)
	if err != nil {
		panic(err)
	}

	addHeaders(req, d.Endpoint.Headers)
	req.Header.Set("Content-Type", "application/json")

	return req
}
func (d *Destination) Fetch() FetchResult {
	timeStart := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var req *http.Request
	if d.Endpoint.Method == "GET" {
		req = createGetReq(ctx, d)
	} else if d.Endpoint.Method == "POST" {
		req = createPostReq(ctx, d)
	} else {
		panic("Method not supported")
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return ErrorResult
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrorResult
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ErrorResult
	}

	return FetchResult{
		ReplyTime:       int(time.Since(timeStart).Milliseconds()),
		Id:              d.Id,
		StatusCode:      resp.StatusCode,
		Body:            string(bodyBytes),
		ResponseHeaders: resp.Header,
	}
}

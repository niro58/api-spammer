package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Destination struct {
	Id     int                    `json:"id"`
	Url    string                 `json:"url"`
	Method string                 `json:"method"`
	Data   map[string]interface{} `json:"data"`
}
type FetchResult struct {
	Id              int
	StatusCode      int
	ResponseHeaders http.Header
	Body            string
	ReplyTime       int
}

var ErrorResult = FetchResult{StatusCode: 0, Body: "", ReplyTime: 0}

func CreateGetReq(ctx context.Context, d *Destination) *http.Request {
	query := url.Values{}
	for k, v := range d.Data {
		query.Add(k, v.(string))
	}
	baseUrl := d.Url
	fullUrl := baseUrl + "?" + query.Encode()

	req, err := http.NewRequestWithContext(ctx, d.Method, fullUrl, nil)
	if err != nil {
		panic(err)
	}
	return req
}
func CreatePostReq(ctx context.Context, d *Destination) *http.Request {
	reqBytes, err := json.Marshal(d.Data)
	if err != nil {
		panic(err)
	}
	reqBody := bytes.NewBuffer(reqBytes)

	req, err := http.NewRequestWithContext(ctx, d.Method, d.Url, reqBody)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	return req
}
func (d *Destination) Fetch() FetchResult {
	timeStart := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var req *http.Request
	if d.Method == "GET" {
		req = CreateGetReq(ctx, d)
	} else if d.Method == "POST" {
		req = CreatePostReq(ctx, d)
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

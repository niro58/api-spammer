package main

import (
	Logger "api-spammer/logger"
)

var config Config
var statistics Statistics

func worker(jobs <-chan Destination, results chan<- FetchResult) {
	for j := range jobs {
		results <- j.Fetch()
	}
}

func main() {
	config = LoadConfig()
	Logger.Init()

	jobs := make(chan Destination, config.TotalRequests)
	results := make(chan FetchResult, config.TotalRequests)

	for w := 0; w < config.Clients; w++ {
		go worker(jobs, results)
	}
	for j := 0; j < config.TotalRequests; j++ {

		ep := config.Endpoints[j%len(config.Endpoints)]

		job := Destination{
			Id:     j,
			Url:    ep.Url,
			Method: ep.Method,
			Data:   ep.Data,
		}

		jobs <- job
	}
	close(jobs)

	for a := 0; a < config.TotalRequests; a++ {
		res := <-results
		statistics.AddRequest(res)
	}

	statistics.Debug()
}

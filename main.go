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

	for w := 1; w <= config.Clients; w++ {
		go worker(jobs, results)
	}
	for j := 1; j <= config.TotalRequests; j++ {
		ep := config.Endpoints[0]
		job := Destination{
			Id:     j,
			Url:    ep.Url,
			Method: ep.Method,
			Data:   ep.Data,
		}

		jobs <- job
	}
	close(jobs)

	for a := 1; a <= config.TotalRequests; a++ {
		res := <-results
		statistics.AddRequest(res)
	}

	statistics.Debug()
}

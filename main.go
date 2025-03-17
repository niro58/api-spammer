package main

import (
	cfg "api-spammer/internal/config"
	"api-spammer/internal/fetcher"
	"api-spammer/internal/logger"
)

var config cfg.Config
var statistics fetcher.Statistics

func worker(jobs <-chan fetcher.Destination, results chan<- fetcher.FetchResult) {
	for j := range jobs {
		results <- j.Fetch()
	}
}

func main() {
	config = cfg.LoadConfig()
	logger.Init()

	jobs := make(chan fetcher.Destination, config.TotalRequests)
	results := make(chan fetcher.FetchResult, config.TotalRequests)

	for w := 0; w < config.Clients; w++ {
		go worker(jobs, results)
	}
	for j := 0; j < config.TotalRequests; j++ {

		ep := config.Endpoints[j%len(config.Endpoints)]

		job := fetcher.Destination{
			Id:       j,
			Endpoint: ep,
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

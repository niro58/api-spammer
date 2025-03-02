package main

import (
	Logger "api-spammer/logger"
	"fmt"
)

type Statistics struct {
	TotalRequests      int
	SuccessfulRequests int
	FailedRequests     int
	TotalTime          int
	MinTime            int
	MaxTime            int
	AverageTime        int
}

func (s *Statistics) AddRequest(result FetchResult) {
	Logger.WriteLog(result)

	s.TotalRequests++
	s.TotalTime += result.ReplyTime

	if result.StatusCode == 200 {
		s.SuccessfulRequests++
	} else {
		s.FailedRequests++
	}

	if s.MinTime == 0 || result.ReplyTime < s.MinTime {
		s.MinTime = result.ReplyTime
	}

	if result.ReplyTime > s.MaxTime {
		s.MaxTime = result.ReplyTime
	}

	s.AverageTime = s.TotalTime / s.TotalRequests
}

func (s *Statistics) Debug() {
	Logger.Print("info", "Total requests: ", s.TotalRequests)
	Logger.Print("info", "Successful requests: ", s.SuccessfulRequests)
	Logger.Print("info", "Failed requests: ", s.FailedRequests)
	Logger.Print("info", fmt.Sprintf("Total time: %d ms", s.TotalTime))
	Logger.Print("info", fmt.Sprintf("Min time: %d ms", s.MinTime))
	Logger.Print("info", fmt.Sprintf("Max time: %d ms", s.MaxTime))
	Logger.Print("info", fmt.Sprintf("Average time: %d ms", s.AverageTime))
}

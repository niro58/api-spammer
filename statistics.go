package main

import (
	logger "api-spammer/logger"
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
	logger.WriteLog(result)

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
	logger.Log(logger.ColorSuccess, "Total requests: ", s.TotalRequests)
	logger.Log(logger.ColorSuccess, "Successful requests: ", s.SuccessfulRequests)
	logger.Log(logger.ColorError, "Failed requests: ", s.FailedRequests)
	logger.Log(logger.ColorDefault, fmt.Sprintf("Total time: %d ms", s.TotalTime))
	logger.Log(logger.ColorDefault, fmt.Sprintf("Min time: %d ms", s.MinTime))
	logger.Log(logger.ColorError, fmt.Sprintf("Max time: %d ms", s.MaxTime))
	logger.Log(logger.ColorDefault, fmt.Sprintf("Average time: %d ms", s.AverageTime))
}

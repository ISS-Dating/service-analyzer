package service

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/ISS-Dating/service-analyzer/model"
	"github.com/ISS-Dating/service-analyzer/repo"
)

type Service struct {
	Repo repo.Interface
}

func NewService(repo repo.Interface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) UpdateWithMessage(message string) {
	var info model.UserMessage
	json.Unmarshal([]byte(message), &info)

	id, _ := strconv.Atoi(info.Sender)

	stats, _ := s.Repo.ReadUserStatsById(uint64(id))
	stats.AverageMessageLen = uint(upateLength(int(stats.AverageMessageLen), int(stats.MessagesSent), info.Body))
	stats.MessagesSent++
	stats.LinksInMessages += uint(countLinks(info.Body))

	s.Repo.UpdateUserStats(stats)
}

func countLinks(message string) int {
	return strings.Count(message, "http://")
}

func upateLength(old int, count int, message string) int {
	return (old*count + len(message)) / (count + 1)
}

package Events

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type EventStruct struct {
	Events []struct {
		Id             int    `json:"id"`
		ParentId       int    `json:"parentId,omitempty"`
		SortOrder      string `json:"sortOrder"`
		Level          int    `json:"level"`
		Num            int    `json:"num"`
		SportId        int    `json:"sportId"`
		Kind           int    `json:"kind"`
		RootKind       int    `json:"rootKind"`
		Team1Id        int    `json:"team1Id"`
		Team2Id        int    `json:"team2Id"`
		Team1          string `json:"team1"`
		Team2          string `json:"team2"`
		Name           string `json:"name"`
		StartTime      int    `json:"startTime"`
		Place          string `json:"place"`
		StatisticsType string `json:"statisticsType"`
		Priority       int    `json:"priority"`
	} `json:"events"`
}

func (fonbet *EventStruct) Parse(url string, logger *logrus.Logger) error {

	request, err := http.Get(url)
	if err != nil {
		logger.Errorf("Cant Parse URL: %v  error: %v", url, err)
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		logger.Errorf("Cant ReadALL URL: %v  error: %v", url, err)

	}
	err = json.Unmarshal(body, &fonbet)
	if err != nil {
		logger.Errorf("Cant Unmarshall URL: %v  error: %v", url, err)
	}
	err = request.Body.Close()
	if err != nil {
		logger.Errorf("Unable to close body URL: %v  error: %v", url, err)
	}
	return err
}

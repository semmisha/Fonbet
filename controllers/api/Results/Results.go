package ApiResults

import (
	"Fonbet/controllers/api"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
)

type ApiResults struct {
	Events []struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Team1Score int
		Team2Score int
		Score      string `json:"score"`
		StartTime  int    `json:"startTime"`
		Status     int    `json:"status"`
	} `json:"events"`
	Sections []struct {
		Id                  int    `json:"id"`
		Events              []int  `json:"events"`
		FonbetCompetitionId int    `json:"fonbetCompetitionId"`
		Name                string `json:"name"`
	} `json:"sections"`
	LineDate int64 `json:"lineDate"`
}

func NewApiResults() *ApiResults {
	return &ApiResults{}
}

func (fonbet *ApiResults) JsonToStruct(url *api.ListURLStruct, logger *logrus.Logger) {

	response := api.Carousele(url.Common, "/results/results.json.php?locale=ru")
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("Cant ReadALL URL: %v  error: %v", url, err)

	}
	err = json.Unmarshal(body, &fonbet)
	if err != nil {
		logger.Errorf("Cant Unmarshall URL: %v  error: %v", url, err)
	}
	err = response.Body.Close()
	if err != nil {
		logger.Errorf("Unable to close body URL: %v  error: %v", url, err)
	}

}

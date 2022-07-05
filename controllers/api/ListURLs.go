package api

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type ListURLStruct struct {
	Common           []string `json:"common"`
	CommonOther      string   `json:"commonOther"`
	Line             []string `json:"line"`
	LineDesktopMaLja []string `json:"line.desktop-ma-lja"`
	LineDesktop      []string `json:"line.desktop"`
	LineLandings     []string `json:"line.landings"`
	StaticN          string   `json:"staticN"`
	OriginN          string   `json:"originN"`
	StaticV          string   `json:"staticV"`
	OriginV          string   `json:"originV"`
	Static           string   `json:"static"`
	Origin           string   `json:"origin"`
	Shop             string   `json:"shop"`
	Statistic        string   `json:"statistic"`
}

func NewListURLStruct() *ListURLStruct {
	return &ListURLStruct{}
}

func (fonbet *ListURLStruct) JsonToStruct(url string, logger *logrus.Logger) error {
	for i := 0; i < 5; i++ {
		request, err := http.Get(url)
		if err != nil {
			logger.Errorf("Cant JsonToStruct URL: %v  error: %v", url, err)
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
		if err == nil {
			return nil
		}
		time.Sleep(3 * time.Minute)
	}
	return errors.New("Failed to parse list of URLs")
}

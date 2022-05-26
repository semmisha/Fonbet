package Factors

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type CustomFactorsStruct struct {
	CustomFactors []struct {
		E        int `json:"e"`
		CountAll int `json:"countAll"`
		Factors  []struct {
			F  int     `json:"f"`
			V  float64 `json:"v"`
			P  int     `json:"p,omitempty"`
			Pt string  `json:"pt,omitempty"`
		} `json:"factors"`
	} `json:"customFactors"`
}

func (fonbet *CustomFactorsStruct) Parse(url string, logger *logrus.Logger) error {

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

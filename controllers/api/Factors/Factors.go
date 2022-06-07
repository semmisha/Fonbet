package ApiFactors

import (
	"Fonbet/controllers/api"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
)

type ApiFactors struct {
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

func (fonbet *ApiFactors) JsonToStruct(url *api.ListURLStruct, logger *logrus.Logger) error {

	response := api.Carousele(url.LineDesktop, "/events/list")

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
	return err
}

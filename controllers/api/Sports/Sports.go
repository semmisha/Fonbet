package ApiSports

import (
	"Fonbet/controllers/api"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
)

type ApiSports struct {
	Sports []struct {
		Id        int    `json:"id"`
		Kind      string `json:"kind"`
		SortOrder string `json:"sortOrder"`
		Name      string `json:"name"`
		ParentId  int    `json:"parentId,omitempty"`
		ParentIds []int  `json:"parentIds,omitempty"`
		RegionId  int    `json:"regionId,omitempty"`
	} `json:"sports"`
}

func (fonbet *ApiSports) JsonToStruct(url *api.ListURLStruct, logger *logrus.Logger) error {

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

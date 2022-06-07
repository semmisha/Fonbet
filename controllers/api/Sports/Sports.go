package Sports

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
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

func (fonbet *ApiSports) JsonToStruct(url string, logger *logrus.Logger) error {

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
	return err
}

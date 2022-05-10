package api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func Parse[T any](fonbet *T, url string, logger *logrus.Logger) {

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

}

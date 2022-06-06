package api

import (
	"fmt"
	"net/http"
)

func Carousele(url []string, extension string) *http.Response {
	for _, i := range url {
		i = fmt.Sprint(i, extension)
		request, err := http.Get(i)
		if err == nil {
			return request
		} else if err != nil && i == url[(len(url)-1)] {
			panic(err)
		}
	}
	return nil
}

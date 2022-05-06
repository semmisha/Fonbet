package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Parse[T any](fonbet *T, url string) {

	request, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: ", err)
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)

	}
	err = json.Unmarshal(body, &fonbet)
	if err != nil {
		fmt.Println(err)
	}
	err = request.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

}

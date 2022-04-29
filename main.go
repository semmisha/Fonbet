package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	request, err := http.Get("https://line06w.bkfon-resources.com/events/list?lang=ru&version=7987900598&scopeMarket=1600")
	if err != nil {
		fmt.Printf("Error: ", err)
	}

	body, err := io.ReadAll(request.Body)

	json.Unmarshal(body)
	request.Body.Close()
	fmt.Println(string(body))

}

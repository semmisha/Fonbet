package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Parse2() *Fonbet2 {
	var fonbet *Fonbet2
	request, err := http.Get("https://line06w.bkfon-resources.com/events/list?lang=ru&version=7987900598&scopeMarket=1600")
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
	return fonbet
}

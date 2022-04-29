package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	DbConnect()

	var fonbet *Fonbet
	request, err := http.Get("https://line55w.bkfon-resources.com/line/desktop/topEvents3?place=live&sysId=1&lang=ru")
	if err != nil {
		fmt.Printf("Error: ", err)
	}

	body, err := io.ReadAll(request.Body)

	json.Unmarshal(body, &fonbet)
	request.Body.Close()
	//fmt.Printf("%+v", fonbet)

}

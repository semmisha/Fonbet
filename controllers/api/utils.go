package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func Carousele(url []string, extension string) *http.Response {

	for _, i := range url {
		i = fmt.Sprint("https:", i, extension)
		//fmt.Println(i)
		request, err := http.Get(i)
		if err == nil {
			return request
		} else if err != nil && i == url[(len(url)-1)] {
			panic(err)
		}
	}
	return nil
}

func RandomNum(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}

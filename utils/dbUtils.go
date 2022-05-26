package utils

import (
	"fmt"
	"strings"
	"time"
)

func DayCount(url string, day int) (urldate string) {

	year, month, day := time.Now().AddDate(0, 0, -day).Date()
	urldate = fmt.Sprintf("%v&lineDate=%v-%02v-%02v", url, year, int(month), day)
	//fmt.Println(urldate)
	return
}

func Replacer(str string) string {
	var symbols = [...]string{" ", "-", "–", "(", ")"}
	for _, i := range symbols {
		str = strings.ReplaceAll(str, i, "")
	}

	return str
}

package utils

import (
	fonstruct "Fonbet/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

func DayCount(url string, day int) (urldate string) {

	year, month, day := time.Now().AddDate(0, 0, -day).Date()
	urldate = fmt.Sprintf("%v&lineDate=%v-%02v-%02v", url, year, int(month), day)
	//fmt.Println(urldate)
	return
}

func SearchSportId(fonbet *fonstruct.FonbetResult, i int, logger *logrus.Logger) int {

	for j := 0; j < len(fonbet.Sections); j++ {
		for b := 0; b < len(fonbet.Sections[j].Events); b++ {
			d, err := strconv.Atoi(fonbet.Events[i].Id)
			if err != nil {
				logger.Warningf("Unable to convert string to int: %v  error:%v\n", d, err)
			}
			if fonbet.Sections[j].Events[b] == d {
				return fonbet.Sections[j].FonbetCompetitionId
			}
		}

	}

	logger.Errorf("Cant find sport id for: %v\n ", fonbet.Events[i].Name)
	return 0
}

func Replacer(str string) string {
	var symbols = [...]string{" ", "-", "–", "(", ")"}
	for _, i := range symbols {
		str = strings.ReplaceAll(str, i, "")
	}

	return str
}

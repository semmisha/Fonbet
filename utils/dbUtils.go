package utils

import (
	fonstruct "Fonbet/json"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func CheckLevels(fonbet *fonstruct.FonbetEvents, levels *int) {
	for i := 0; i < len(fonbet.Events); i++ {

		if *levels < fonbet.Events[i].Level {
			*levels = fonbet.Events[i].Level
		}

	}

}
func CreateLevels(db *pgxpool.Pool, levels *int) {

	for i := 2; i <= *levels; i++ {
		currentlevel := fmt.Sprintf("events_level_%v", i)
		parentlevel := fmt.Sprintf("events_level_%v", i-1)

		query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (id INT PRIMARY KEY, parentid int, eventname VARCHAR(150), sportid INT, team1 VARCHAR(50), team2 VARCHAR(50), starttime INT, foreign key (parentid) references %v (id) )", currentlevel, parentlevel)
		_, err := db.Exec(context.Background(), query)
		if err != nil {
			fmt.Println(err)
		}

	}

}

func DayCount(url string, day int) (urldate string) {

	year, month, day := time.Now().AddDate(0, 0, -day).Date()
	urldate = fmt.Sprintf("%v&lineDate=%v-%02v-%02v", url, year, int(month), day)
	//fmt.Println(urldate)
	return
}

func SearchSportId(fonbet *fonstruct.FonbetResult, i int, logger *logrus.Logger) int {

	for j := 0; j < len(fonbet.Sections); j++ {
		for b := 0; b < len(fonbet.Sections[j].Events); b++ {
			d, _ := strconv.Atoi(fonbet.Events[i].Id)
			if fonbet.Sections[j].Events[b] == d {
				return fonbet.Sections[j].FonbetCompetitionId
			}
		}

	}

	logger.Errorf("Cant find sport id for: %v\n ", fonbet.Events[i].Name)
	return 0
}

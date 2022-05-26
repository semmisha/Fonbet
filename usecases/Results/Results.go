package Results

import (
	"Fonbet/controllers/api/Results"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type UcResultSlice struct {
	UcResultsStruct []UcResultsStruct
}

type UcResultsStruct struct {
	eventid    int
	id         int
	name       string
	team1Score int
	team2Score int
	totalScore string
	startTime  time.Time
	sportId    int
}

func (f *UcResultSlice) ReAssign(fonbet Results.ResultsStruct, logger *logrus.Logger) {

	for i := 0; i < len(fonbet.Events); i++ {
		stringId, err := strconv.Atoi(fonbet.Events[i].Id)
		if err != nil {
			logger.Errorf("Cant strconv Atoi, id %v   error: %v\n", fonbet.Events[i].Id, err)
		}

		for j := 0; j < len(fonbet.Sections); j++ {
			for b := 0; b < len(fonbet.Sections[j].Events); b++ {

				if fonbet.Sections[j].Events[b] == stringId && strings.ContainsAny(fonbet.Events[i].Name, "-––") && fonbet.Events[i].Status == 3 && strings.Contains(fonbet.Events[i].Name, "очковые") != true {

					resultslice, err := Conv(fonbet.Events[i].Score)
					if err != nil {
						logger.Errorf("Unable to convert Result, error:", err)

					}

					fontime := time.Unix(int64(fonbet.Events[i].StartTime), 0)
					g := UcResultsStruct{
						id:         stringId,
						name:       fonbet.Events[i].Name,
						team1Score: resultslice[0],
						team2Score: resultslice[1],
						totalScore: fonbet.Events[i].Score,
						startTime:  fontime,
						sportId:    fonbet.Sections[j].FonbetCompetitionId,
					}

					f.UcResultsStruct = append(f.UcResultsStruct, g)

				}
			}
		}
	}
	//fmt.Println(f)
}

func Conv(fonbet string) ([]int, error) {

	strarray := strings.Split(fonbet, " ")

	strarray = strings.Split(strarray[0], ":")

	resultslice := make([]int, len(strarray))

	if len(resultslice) >= 2 && strarray[0] != "" && strarray[1] != "" {
		for i := 0; i < len(strarray); i++ {
			resultslice[i], _ = strconv.Atoi(strarray[i])

		}

		return resultslice, nil
	}
	return nil, nil
}

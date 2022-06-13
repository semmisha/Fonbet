package Convert

import (
	ApiResults "Fonbet/controllers/api/Results"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type UcResults struct {
	UcResultsStruct []Result
}

type Result struct {
	ResultId   int
	EventId    int
	Name       string
	Team1Score int
	Team2Score int
	TotalScore string
	StartTime  time.Time
	SportId    int
}

func (f *UcResults) ReAssign(fonbet ApiResults.ApiResults, logger *logrus.Logger) {

	for i := 0; i < len(fonbet.Events); i++ {
		stringId, err := strconv.Atoi(fonbet.Events[i].Id)
		if err != nil {
			logger.Errorf("Cant strconv Atoi, ResultId %v   error: %v\n", fonbet.Events[i].Id, err)
		}

		for j := 0; j < len(fonbet.Sections); j++ {
			for b := 0; b < len(fonbet.Sections[j].Events); b++ {

				if fonbet.Sections[j].Events[b] == stringId && strings.ContainsAny(fonbet.Events[i].Name, "-––") && fonbet.Events[i].Status == 3 && strings.Contains(fonbet.Events[i].Name, "очковые") && (fonbet.Sections[j].FonbetCompetitionId < 71000 || fonbet.Sections[j].FonbetCompetitionId > 79000) {

					resultslice, err := Conv(fonbet.Events[i].Score)
					if err != nil {
						logger.Errorf("Unable to convert Result, error:", err)

					}

					fontime := time.Unix(int64(fonbet.Events[i].StartTime), 0)

					g := Result{
						ResultId:   stringId,
						Name:       fonbet.Events[i].Name,
						Team1Score: resultslice[0],
						Team2Score: resultslice[1],
						TotalScore: fonbet.Events[i].Score,
						StartTime:  fontime,
						SportId:    fonbet.Sections[j].FonbetCompetitionId,
					}

					f.UcResultsStruct = append(f.UcResultsStruct, g)

				}
			}
		}
	}

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

package Compare

import (
	. "Fonbet/repository/postgres/Manipulate"
	"Fonbet/utils"
	"fmt"
	"github.com/sirupsen/logrus"
)

func CompareResult(events DbEvents, result DbResults, logger *logrus.Logger) DbResults {
	var (
		count        = 0
		returnResult DbResults
	)
	for _, i := range result.UcResultsStruct {

		resultString := utils.Replacer(i.Name)

		for _, j := range events.UcEventStruct {

			eventString := utils.Replacer(fmt.Sprint(j.Team1, j.Team2))
			if resultString == eventString && i.SportId == j.SportId {
				i.EventId = j.Id

				returnResult.UcResultsStruct = append(returnResult.UcResultsStruct, i)

				count++

			} else {

				//fmt.Println(resultString, eventString)

			}

		}

	}
	logger.Infof("New copmare entries: %v\n", count)
	return returnResult
}

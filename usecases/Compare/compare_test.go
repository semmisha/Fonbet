package Compare

import (
	"Fonbet/logging"
	Postgres "Fonbet/repository/postgres/Manipulate"
	UcEvents "Fonbet/usecases/Convert"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
	"time"
)

func TestCompareResult(t *testing.T) {
	type args struct {
		events  Postgres.DbEvents
		results Postgres.DbResults
		want    Postgres.DbResults
		logger  *logrus.Logger
	}
	tests := args{
		events: Postgres.DbEvents{
			UcEventStruct: []UcEvents.Event{{
				Id:        34741123,
				SportId:   11676,
				Team1Id:   551095,
				Team2Id:   440754,
				Team1:     "Детройт Тайгерс",
				Team2:     "Торонто Блю Джейс",
				Name:      "",
				StartTime: time.Time{},
			}},
		},
		results: Postgres.DbResults{
			UcResultsStruct: []UcEvents.Result{{
				ResultId:   7991,
				EventId:    0,
				Name:       "Детройт Тайгерс – Торонто Блю Джейс",
				Team1Score: 0,
				Team2Score: 6,
				TotalScore: "0:6 (0-0 0-0 0-0 0-4 0-0 0-0 0-0 0-2 0-0)",
				StartTime:  time.Time{},
				SportId:    11676,
			}},
		},
		want: Postgres.DbResults{
			UcResultsStruct: []UcEvents.Result{{
				ResultId:   7991,
				EventId:    34741123,
				Name:       "Детройт Тайгерс – Торонто Блю Джейс",
				Team1Score: 0,
				Team2Score: 6,
				TotalScore: "0:6 (0-0 0-0 0-0 0-4 0-0 0-0 0-0 0-2 0-0)",
				StartTime:  time.Time{},
				SportId:    11676,
			}},
		},
		logger: logging.Logger(),
	}

	t.Run("Compare", func(t *testing.T) {
		if got := CompareResult(tests.events, tests.results, tests.logger); !reflect.DeepEqual(got, tests.want) {
			t.Errorf("CompareResult() = %v, want %v", got, tests.want)
		}
	})

}

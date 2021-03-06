package Convert

import (
	ApiEvents "Fonbet/controllers/api/Events"
	"time"
)

type UcEvents struct {
	UcEventStruct []Event
}

func NewUcEvents() *UcEvents {
	return &UcEvents{}
}

type Event struct {
	Id        int
	SportId   int
	Team1Id   int
	Team2Id   int
	Team1     string
	Team2     string
	Name      string
	StartTime time.Time
}

func (f *UcEvents) ReAssign(fonbet ApiEvents.ApiEvents) {
	for i := 0; i < len(fonbet.Events); i++ {
		if fonbet.Events[i].Team2Id != 0 && fonbet.Events[i].Team1Id != 0 && fonbet.Events[i].Level == 1 {

			fontime := time.Unix(int64(fonbet.Events[i].StartTime), 0)
			j := Event{
				Id:        fonbet.Events[i].Id,
				SportId:   fonbet.Events[i].SportId,
				Team1Id:   fonbet.Events[i].Team1Id,
				Team2Id:   fonbet.Events[i].Team2Id,
				Team1:     fonbet.Events[i].Team1,
				Team2:     fonbet.Events[i].Team2,
				Name:      fonbet.Events[i].Name,
				StartTime: fontime,
			}
			f.UcEventStruct = append(f.UcEventStruct, j)

		}

	}
	//fmt.Println(f)
}

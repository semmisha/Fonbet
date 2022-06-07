package Sports

import (
	"Fonbet/controllers/api/Sports"
)

type UcSports struct {
	UcSportsStruct []Sport
}

type Sport struct {
	id       int
	parentid int
	name     string
}

func (f *UcSports) ReAssign(fonbet Sports.ApiSports) {
	for i := 0; i < len(fonbet.Sports); i++ {
		j := Sport{
			id:       fonbet.Sports[i].Id,
			parentid: fonbet.Sports[i].ParentId,
			name:     fonbet.Sports[i].Name,
		}
		f.UcSportsStruct = append(f.UcSportsStruct, j)

	}
	//fmt.Println(f)
}

package Sports

import (
	"Fonbet/controllers/api/Sports"
)

type UcSportsSlice struct {
	UcSportsStruct []UcSportsStruct
}

type UcSportsStruct struct {
	id       int
	parentid int
	name     string
}

func (f *UcSportsSlice) ReAssign(fonbet Sports.SportsStruct) {
	for i := 0; i < len(fonbet.Sports); i++ {
		j := UcSportsStruct{
			id:       fonbet.Sports[i].Id,
			parentid: fonbet.Sports[i].ParentId,
			name:     fonbet.Sports[i].Name,
		}
		f.UcSportsStruct = append(f.UcSportsStruct, j)

	}
	//fmt.Println(f)
}

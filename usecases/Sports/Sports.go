package UcSports

import ApiSports "Fonbet/controllers/api/Sports"

type UcSports struct {
	UcSportsStruct []Sport
}

type Sport struct {
	Id       int
	ParentId int
	Name     string
}

func (f *UcSports) ReAssign(fonbet ApiSports.ApiSports) {
	for i := 0; i < len(fonbet.Sports); i++ {
		j := Sport{
			Id:       fonbet.Sports[i].Id,
			ParentId: fonbet.Sports[i].ParentId,
			Name:     fonbet.Sports[i].Name,
		}
		f.UcSportsStruct = append(f.UcSportsStruct, j)

	}
	//fmt.Println(f)
}

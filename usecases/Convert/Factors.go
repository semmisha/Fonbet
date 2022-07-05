package Convert

import ApiFactors "Fonbet/controllers/api/Factors"

type UcFactors struct {
	UcFactorsStruct []Factor
}

func NewUcFactors() *UcFactors {
	return &UcFactors{}
}

type Factor struct {
	Id     int
	FrstWn float64
	Drw    float64
	ScndWn float64
}

func (f *UcFactors) ReAssign(fonbet ApiFactors.ApiFactors) {

	for i := 0; i < len(fonbet.CustomFactors); i++ {
		g := Factor{
			Id:     fonbet.CustomFactors[i].E,
			FrstWn: 0,
			Drw:    0,
			ScndWn: 0,
		}
		for j := 0; j < len(fonbet.CustomFactors[i].Factors); j++ {
			switch fonbet.CustomFactors[i].Factors[j].F {
			case 921:
				g.FrstWn = fonbet.CustomFactors[i].Factors[j].V
			case 922:
				g.Drw = fonbet.CustomFactors[i].Factors[j].V
			case 923:
				g.ScndWn = fonbet.CustomFactors[i].Factors[j].V
			}
		}
		if g.Id != 0 && (g.FrstWn != 0 || g.Drw != 0 || g.ScndWn != 0) {
			f.UcFactorsStruct = append(f.UcFactorsStruct, g)
		}
	}
	//fmt.Println(f)
}

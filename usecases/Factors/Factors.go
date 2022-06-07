package Factors

import (
	"Fonbet/controllers/api/Factors"
)

type UcFactors struct {
	UcFactorsStruct []Factor
}

type Factor struct {
	id     int
	frstWn float64
	drw    float64
	scndWn float64
}

func (f *UcFactors) ReAssign(fonbet Factors.ApiFactors) {

	for i := 0; i < len(fonbet.CustomFactors); i++ {
		g := Factor{
			id:     fonbet.CustomFactors[i].E,
			frstWn: 0,
			drw:    0,
			scndWn: 0,
		}
		for j := 0; j < len(fonbet.CustomFactors[i].Factors); j++ {
			switch fonbet.CustomFactors[i].Factors[j].F {
			case 921:
				g.frstWn = fonbet.CustomFactors[i].Factors[j].V
			case 922:
				g.drw = fonbet.CustomFactors[i].Factors[j].V
			case 923:
				g.scndWn = fonbet.CustomFactors[i].Factors[j].V
			}
		}
		if g.id != 0 && (g.frstWn != 0 || g.drw != 0 || g.scndWn != 0) {
			f.UcFactorsStruct = append(f.UcFactorsStruct, g)
		}
	}
	//fmt.Println(f)
}

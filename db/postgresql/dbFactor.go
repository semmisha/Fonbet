package Postgres

import (
	fonstruct "Fonbet/json"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func Factors(fonbet *fonstruct.FonbetEvents, i int, db *pgxpool.Pool, logger *logrus.Logger) {
	type Factors struct {
		id        int
		firstwin  float64
		draw      float64
		secondwin float64
	}
	var factor = Factors{fonbet.Events[i].Id, 0, 0, 0}

	for _, b := range fonbet.CustomFactors {
		for _, c := range b.Factors {

			switch {

			case b.E == fonbet.Events[i].Id && c.F == 921:
				factor.firstwin = c.V
			case b.E == fonbet.Events[i].Id && c.F == 922:
				factor.draw = c.V
			case b.E == fonbet.Events[i].Id && c.F == 923:
				factor.secondwin = c.V

			}

		}

	}

	query := fmt.Sprintf(`INSERT INTO factors (eventid, "921", "922", "923") Values ( %v, %v , %v, %v )`, factor.id, factor.firstwin, factor.draw, factor.secondwin)
	_, err := db.Exec(context.Background(), query)
	if err != nil {
		logger.Warningf("Unable to insert factors: %v error: %v", factor.id, err)
	}

}

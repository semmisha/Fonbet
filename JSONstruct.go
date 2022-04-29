package main

func init() {

}

type Fonbet struct {
	Events []struct {
		Id                       int    `json:"id"`
		Number                   int    `json:"number"`
		StartTimeTimestamp       int    `json:"startTimeTimestamp"`
		CompetitionId            int    `json:"competitionId"`
		CompetitionsGroupCaption string `json:"competitionsGroupCaption,omitempty"`
		Team1Id                  int    `json:"team1Id"`
		Team2Id                  int    `json:"team2Id"`
		Team1                    string `json:"team1"`
		Team2                    string `json:"team2"`
		Timer                    string `json:"timer,omitempty"`
		TimerSeconds             int    `json:"timerSeconds,omitempty"`
		TimerDirection           int    `json:"timerDirection,omitempty"`
		TimerTimestampMsec       int64  `json:"timerTimestampMsec,omitempty"`
		ScoreFunction            string `json:"scoreFunction"`
		AllFactorsCount          int    `json:"allFactorsCount"`
		Markets                  []struct {
			Rows []struct {
				IsTitle bool `json:"isTitle,omitempty"`
				Cells   []struct {
					IsTitle      bool    `json:"isTitle,omitempty"`
					Caption      string  `json:"caption,omitempty"`
					CaptionAlign string  `json:"captionAlign,omitempty"`
					FactorId     int     `json:"factorId,omitempty"`
					EventId      int     `json:"eventId,omitempty"`
					Value        float64 `json:"value,omitempty"`
					EventKindId  int     `json:"eventKindId,omitempty"`
					ParamText    string  `json:"paramText,omitempty"`
					Param        int     `json:"param,omitempty"`
					FlexParam    bool    `json:"flexParam,omitempty"`
					Blocked      bool    `json:"blocked,omitempty"`
					ParamHigh    int     `json:"paramHigh,omitempty"`
					ParamLow     int     `json:"paramLow,omitempty"`
				} `json:"cells"`
			} `json:"rows"`
		} `json:"markets"`
	} `json:"events"`
}

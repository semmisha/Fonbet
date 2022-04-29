package main

type Fonbet struct {
}

func init() {

}

type T struct {
	Result  string `json:"result"`
	Request string `json:"request"`
	Place   string `json:"place"`
	Lang    string `json:"lang"`
	Events  []struct {
		Id                       int    `json:"id"`
		Number                   int    `json:"number"`
		StartTimeTimestamp       int    `json:"startTimeTimestamp"`
		CompetitionId            int    `json:"competitionId"`
		CompetitionName          string `json:"competitionName"`
		CompetitionCaption       string `json:"competitionCaption"`
		CompetitionsGroupId      string `json:"competitionsGroupId,omitempty"`
		CompetitionsGroupCaption string `json:"competitionsGroupCaption,omitempty"`
		SkId                     int    `json:"skId"`
		SkName                   string `json:"skName"`
		SkSortOrder              string `json:"skSortOrder"`
		RegionId                 int    `json:"regionId,omitempty"`
		Team1Id                  int    `json:"team1Id"`
		Team2Id                  int    `json:"team2Id"`
		Team1                    string `json:"team1"`
		Team2                    string `json:"team2"`
		StatisticsType           string `json:"statisticsType"`
		EventName                string `json:"eventName"`
		Name                     string `json:"name"`
		Place                    string `json:"place"`
		Priority                 int    `json:"priority"`
		Kind                     int    `json:"kind"`
		RootKind                 int    `json:"rootKind"`
		SortOrder                string `json:"sortOrder"`
		Tv                       []int  `json:"tv"`
		SportViewId              int    `json:"sportViewId"`
		Timer                    string `json:"timer,omitempty"`
		TimerSeconds             int    `json:"timerSeconds,omitempty"`
		TimerDirection           int    `json:"timerDirection,omitempty"`
		TimerTimestampMsec       int64  `json:"timerTimestampMsec,omitempty"`
		ScoreFunction            string `json:"scoreFunction"`
		Scores                   [][]struct {
			C1      string `json:"c1"`
			C2      string `json:"c2"`
			Title   string `json:"title,omitempty"`
			Serve   int    `json:"serve,omitempty"`
			Comment string `json:"comment,omitempty"`
		} `json:"scores"`
		Subscores []struct {
			KindId   string `json:"kindId"`
			KindName string `json:"kindName"`
			C1       string `json:"c1"`
			C2       string `json:"c2"`
			Alias    string `json:"alias,omitempty"`
			Comment  string `json:"comment,omitempty"`
		} `json:"subscores"`
		AllFactorsCount int `json:"allFactorsCount"`
		Markets         []struct {
			MarketId      string        `json:"marketId"`
			Ident         string        `json:"ident"`
			SortOrder     int           `json:"sortOrder"`
			Caption       string        `json:"caption"`
			CommonHeaders []interface{} `json:"commonHeaders"`
			Rows          []struct {
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
		ScoreComment     string `json:"scoreComment,omitempty"`
		WillBeLive       bool   `json:"willBeLive,omitempty"`
		ScoreCommentTail string `json:"scoreCommentTail,omitempty"`
	} `json:"events"`
	Md5 string `json:"md5"`
}

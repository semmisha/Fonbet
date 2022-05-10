package json

type FonbetResult struct {
	Events []struct {
		Comment1  string `json:"comment1"`
		Comment2  string `json:"comment2"`
		Comment3  string `json:"comment3"`
		GoalOrder string `json:"goalOrder"`
		Id        string `json:"id"`
		Name      string `json:"name"`
		Score     string `json:"score"`
		StartTime int64  `json:"startTime"`
		Status    int    `json:"status"`
	} `json:"events"`
	Sports []struct {
		Name      string `json:"name"`
		FonbetId  int    `json:"fonbetId"`
		SortOrder string `json:"sortOrder"`
		Id        string `json:"id"`
	} `json:"sports"`
	Sections []struct {
		Id                  int    `json:"id"`
		Events              []int  `json:"events"`
		FonbetSportId       int    `json:"fonbetSportId"`
		FonbetCompetitionId int    `json:"fonbetCompetitionId"`
		SortOrder           string `json:"sortOrder"`
		Name                string `json:"name"`
		Sport               string `json:"sport"`
	} `json:"sections"`
	IsArchive                bool   `json:"isArchive"`
	Lang                     string `json:"lang"`
	Generated                int64  `json:"generated"`
	LastArchiveLineDate      int64  `json:"lastArchiveLineDate"`
	LastChangeTimeFromOffice int64  `json:"lastChangeTimeFromOffice"`
	LineDate                 int64  `json:"lineDate"`
}

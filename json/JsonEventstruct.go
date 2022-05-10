package json

type FonbetEvents struct {
	PacketVersion               int64 `json:"packetVersion"`
	FromVersion                 int   `json:"fromVersion"`
	CatalogTablesVersion        int   `json:"catalogTablesVersion"`
	CatalogSpecialTablesVersion int   `json:"catalogSpecialTablesVersion"`
	CatalogEventViewVersion     int   `json:"catalogEventViewVersion"`
	SportBasicMarketsVersion    int   `json:"sportBasicMarketsVersion"`
	SportKindsVersion           int   `json:"sportKindsVersion"`
	TopCompetitionsVersion      int   `json:"topCompetitionsVersion"`
	Sports                      []struct {
		Id        int    `json:"id"`
		Kind      string `json:"kind"`
		SortOrder string `json:"sortOrder"`
		Name      string `json:"name"`
		ParentId  int    `json:"parentId,omitempty"`
		ParentIds []int  `json:"parentIds,omitempty"`
		RegionId  int    `json:"regionId,omitempty"`
	} `json:"sports"`
	Events []struct {
		Id             int    `json:"id"`
		ParentId       int    `json:"parentId,omitempty"`
		SortOrder      string `json:"sortOrder"`
		Level          int    `json:"level"`
		Num            int    `json:"num"`
		SportId        int    `json:"sportId"`
		Kind           int    `json:"kind"`
		RootKind       int    `json:"rootKind"`
		Team1Id        int    `json:"team1Id"`
		Team2Id        int    `json:"team2Id"`
		Team1          string `json:"team1"`
		Team2          string `json:"team2"`
		Name           string `json:"name"`
		StartTime      int64  `json:"startTime"`
		Place          string `json:"place"`
		StatisticsType string `json:"statisticsType"`
		Priority       int    `json:"priority"`
	} `json:"events"`
	CustomFactors []struct {
		E        int `json:"e"`
		CountAll int `json:"countAll"`
		Factors  []struct {
			F  int     `json:"f"`
			V  float64 `json:"v"`
			P  int     `json:"p,omitempty"`
			Pt string  `json:"pt,omitempty"`
		} `json:"factors"`
	} `json:"customFactors"`
}

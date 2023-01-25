package fpl/leader

type LeagueInfo struct {	
	Standings struct {
		HasNext bool          `json:"has_next"`
		Page    int           `json:"page"`
		Results []struct {
			Entry           int       `json:"entry"`
			EntryName       string    `json:"entry_name"`
			JoinedTime      time.Time `json:"joined_time"`
			PlayerName string    `json:"player_name"`			
		} `json:"results"`
	} `json:"standings"`
}

type ManagerGameweekDetails struct {
	Picks []struct {
		Element	int	`json:"element"`
	}
}

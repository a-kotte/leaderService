package main

import "time"

type EventStatus struct{
	Status					[]struct {
		BonusAdded			bool		`json:"bonus_added"`
		Event				int			`json:"event"`		
	} `json:"status"`	
}

type BootstrapInfo struct{
	Elements				[]struct {
		FirstName			string		`json:"first_name"`
		SecondName			string		`json:"second_name"`
		ID 					int			`json:"id"`	

	} `json:"elements"`
}

type LeagueInfo struct {	
	Standings struct {
		HasNext bool          `json:"has_next"`
		Page    int           `json:"page"`
		Results []struct {
			Entry           int       `json:"entry"`
			EntryName       string    `json:"entry_name"`
			JoinedTime      time.Time `json:"joined_time"`
			PlayerName 		string    `json:"player_name"`			
		} `json:"results"`
	} `json:"standings"`
}

type ManagerGameweekDetails struct {
	Picks []struct {
		Element	int	`json:"element"`
	}
}

type TeamWeekly struct {
	ActiveChip    string        `json:"active_chip"`	
	EntryHistory  struct {
		Event              int `json:"event"`
		Points             int `json:"points"`
		TotalPoints        int `json:"total_points"`
		Rank               int `json:"rank"`
		RankSort           int `json:"rank_sort"`
		OverallRank        int `json:"overall_rank"`
		Bank               int `json:"bank"`
		Value              int `json:"value"`
		EventTransfers     int `json:"event_transfers"`
		EventTransfersCost int `json:"event_transfers_cost"`
		PointsOnBench      int `json:"points_on_bench"`
	} `json:"entry_history"`
	Picks []struct {
		Element       int  `json:"element"`
		Position      int  `json:"position"`
		Multiplier    int  `json:"multiplier"`
		IsCaptain     bool `json:"is_captain"`
		IsViceCaptain bool `json:"is_vice_captain"`
	} `json:"picks"`
}

type TransferHistory[] struct {
	ElementIn      int       `json:"element_in"`
	ElementInCost  int       `json:"element_in_cost"`
	ElementOut     int       `json:"element_out"`
	ElementOutCost int       `json:"element_out_cost"`
	Entry          int       `json:"entry"`	
	Time           time.Time `json:"time"`
	Event          int       `json:"event,omitempty"`
}
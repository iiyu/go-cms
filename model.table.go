package main

type Table struct {
	BaseModel
	CreatedById   uint       `json:"-" gorm:"column:create_user_id"`
	CreatedBy     *UserTable `json:"created_by" gorm:"ForeignKey:CreatedById"`
	Name          string     `json:"table_name"`
	GameMod       string     `json:"-"`
	BigBlind      int        `json:"big_blind"`
	LimitPlayers  int        `json:"limit_players"`
	EndTime       int64      `json:"end_time"`
	ExistTime     int        `json:"exist_time"`
	HandCounts    int        `json:"hand_counts"`
	TotalBuyin    int        `json:"total_buyin"`
	InsurancePool int        `json:"insurance_pool"`
}

type TableExport struct {
	GameMod       string `xlsx:"-"`
	Name          string `xlsx:"table_name"`
	CreateName    string `xlsx:"create_name"`
	BigBlind      string `xlsx:"big_blind"`
	LimitPlayers  int    `xlsx:"limit_players"`
	ExistTime     string `xlsx:"exist_time"`
	HandCounts    int    `xlsx:"hand_counts"`
	Uno           string `xlsx:"u_no"`
	Username      string `xlsx:"0"`
	ClubNo        string `xlsx:"club_no"`
	ClubName      string `xlsx:"club_name"`
	Spends        int    `xlsx:"spends"`
	GetBack       int    `xlsx:"get_back"`
	InsurancePay  int    `xlsx:"insurance_pay"`
	InsuranceBet  int    `xlsx:"insurance_bet"`
	InsuranceNum  int    `xlsx:"insurance_num"`
	InsurancePool int    `xlsx:"insurance_pool"`
	TotalBuyin    int    `xlsx:"total_buyin"`
	Win           int    `xlsx:"win"`
	EndTime       string `xlsx:"end_time"`
}

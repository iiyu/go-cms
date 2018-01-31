package main

type PlayersTables struct {
	BaseModel
	Pid          uint       `json:"-" gorm:"column:pid"`
	User         *UserTable `json:"user" gorm:"ForeignKey:Pid"`
	ClubId       int        `json:"club_id"`
	Club         *Club      `json:"-" gorm:"ForeignKey:ClubId"`
	Tid          int        `json:"tid"`
	Spends       int        `json:"spends"`
	GetBack      int        `json:"get_back"`
	HandCounts   int        `json:"hand_counts"`
	InsuranceNum int        `json:"insurance_num"`
	InsuranceBet int        `json:"insurance_bet"`
	InsurancePay int        `json:"insurance_pay"`
}

package main

type ClubUnion struct {
	BaseModel
	CreatedById uint    `json:"-" gorm:"column:user_id"`
	flag        uint    `json:"-"`
	ClubId      uint    `json:"-" gorm:"column:fid"`
	Clubs       []*Club `json:"clubs"`
}

package main

type Club struct {
	BaseModel
	CreatedById  uint   `json:"-" gorm:"column:user_id"`
	Name         string `json:"club_name"`
	ClubNumber   string `json:"club_no"`
	Funds        uint   `json:"funds"`
	UsersLimit   uint   `json:"users_limit"`
	UsersCount   uint   `json:"users_count"`
	Avatar       string `json:"avatar"`
	Level        int    `json:"level"`
	ManagerLimit int    `json:"manager_limit"`
	ManagerCount int    `json:"manager_count" gorm:"-"`
}

type ClubUser struct {
	BaseModel
	club_id uint `json:"-"`
}

package main

type User struct {
	BaseModel
	Tel      string `json:"tel"`
	Password string `json:"-"`
	Username string `json:"username"`
	Diamonds string `json:"diamonds"`
	Scores   string `json:"scores"`
}
type UserTable struct {
	BaseModel
	Tel      string `json:"tel"`
	Username string `json:"username"`
	Uno      string `json:"u_no" gorm:"column:u_no"`
}

func (UserTable) TableName() string {
	return "players"
}

func (User) TableName() string {
	return "players"
}

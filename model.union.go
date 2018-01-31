package main

type Union struct {
	BaseModel
	CreatedById   uint   `json:"-" gorm:"column:user_id"`
	UnionName     string `json:"union_name"`
	UnionNumber   string `json:"union_no"`
	UnionDescribe string `json:"union_describe"`
	UnionCity     string `json:"union_city"`
	UnionAvatar   string `json:"union_avatar"`
}

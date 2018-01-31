package main

import "github.com/jinzhu/gorm"

type TableRepository interface {
	FindById(id int) ([]*PlayersTables, error)
	FindAll(limit int, offset int) ([]*Table, error)
	FindByUserId(user int) ([]*Table, error)
	FindByClubId(club int, option string, limit int, offset int) ([]*Table, error)
}

type ORMTableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) TableRepository {
	return &ORMTableRepository{db}
}

func (r *ORMTableRepository) FindById(id int) ([]*PlayersTables, error) {
	var playerstables []*PlayersTables

	if err := r.db.Preload("User").Preload("Club").Where("tid = ?", id).Find(&playerstables).Error; err != nil {
		return nil, err
	}

	return playerstables, nil
}

func (r *ORMTableRepository) FindAll(limit int, offset int) ([]*Table, error) {
	var tables []*Table

	err := r.db.Preload("CreatedBy").
		Preload("Tags").
		Limit(limit).
		Offset(offset).
		Find(&tables).Error

	if err != nil {
		return nil, err
	}

	return tables, nil
}

func (r *ORMTableRepository) FindByUserId(user int) ([]*Table, error) {
	var tables []*Table

	err := r.db.Where("create_user_id = ?", user).
		Find(&tables).Error
	if err != nil {
		return nil, err
	}

	return tables, nil
}

func (r *ORMTableRepository) FindByClubId(club int, option string, limit int, offset int) ([]*Table, error) {
	var tables []*Table

	err := r.db.Table("players_tables").Select("`table`.*").Joins("join `table` on `table`.id = players_tables.tid and players_tables.club_id = ? and `table`.is_closed = 0 and "+option, club).
		Preload("CreatedBy").
		Group("`table`.id").
		Order("end_time desc").
		Limit(limit).
		Offset(offset).
		Find(&tables).Error

	if err != nil {
		return nil, err
	}

	return tables, nil
}

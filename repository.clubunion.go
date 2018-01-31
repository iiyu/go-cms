package main

import "github.com/jinzhu/gorm"

type ClubUnionRepository interface {
	FindById(id int) (*Union, error)
	FindAll(limit int, offset int) ([]*Union, error)
	FindByUserId(user int) ([]*Union, error)
	FindClubsByUserId(user int) ([]*Club, error)
	FindByClubId(club int) (int, error)
}

type ORMClubUnionRepository struct {
	db *gorm.DB
}

func NewClubUnionRepository(db *gorm.DB) ClubUnionRepository {
	return &ORMClubUnionRepository{db}
}

func (r *ORMClubUnionRepository) FindById(id int) (*Union, error) {
	club := new(Union)

	if err := r.db.Preload("CreatedBy").Preload("Tags").First(club, id).Error; err != nil {
		return nil, err
	}

	return club, nil
}

func (r *ORMClubUnionRepository) FindAll(limit int, offset int) ([]*Union, error) {
	var unions []*Union

	err := r.db.Preload("CreatedBy").
		Preload("Tags").
		Limit(limit).
		Offset(offset).
		Find(&unions).Error

	if err != nil {
		return nil, err
	}

	return unions, nil
}

func (r *ORMClubUnionRepository) FindByUserId(user int) ([]*Union, error) {
	var unions []*Union

	err := r.db.Where("user_id = ? and flag = '1' and fid = 0", user).
		Find(&unions).Error
	if err != nil {
		return nil, err
	}
	return unions, nil
}

func (r *ORMClubUnionRepository) FindClubsByUserId(user int) ([]*Club, error) {
	var clubs []*Club
	var union Union

	err := r.db.Table("club_union").Where("user_id = ? and flag = '1' and fid = 0", user).First(&union).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Table("club_union").Select("club.id, club.name, club.club_number, club.funds, club.users_limit,  club.users_count, club.avatar").Joins("left join club on club.id = club_union.club_id").Where("flag = '1' and fid = ?", union.ID).Find(&clubs).Error
	if err != nil {
		return nil, err
	}
	return clubs, nil
}

func (r *ORMClubUnionRepository) FindByClubId(club int) (int, error) {
	var unions []int
	err := r.db.Table("club_union").Where("club_id = ? and flag = '1'", club).Pluck("fid", &unions).Error
	if err != nil {
		return 0, err
	}
	return unions[0], nil
}

package main

import "github.com/jinzhu/gorm"

type ClubRepository interface {
	FindById(id int) (*Club, error)
	FindAll(limit int, offset int) ([]*Club, error)
	FindByUserId(user int) ([]*Club, error)
}

type ORMClubRepository struct {
	db *gorm.DB
}

func NewClubRepository(db *gorm.DB) ClubRepository {
	return &ORMClubRepository{db}
}

func (r *ORMClubRepository) FindById(id int) (*Club, error) {
	club := new(Club)

	if err := r.db.Where("status = '1' and id = ?", id).First(club).Error; err != nil {
		return nil, err
	}

	return club, nil
}

func (r *ORMClubRepository) FindAll(limit int, offset int) ([]*Club, error) {
	var clubs []*Club

	err := r.db.Preload("CreatedBy").
		Preload("Tags").
		Limit(limit).
		Offset(offset).
		Find(&clubs).Error

	if err != nil {
		return nil, err
	}

	return clubs, nil
}

func (r *ORMClubRepository) FindByUserId(user int) ([]*Club, error) {
	var clubs []*Club

	err := r.db.Where("user_id = ?", user).
		Find(&clubs).Error
	if err != nil {
		return nil, err
	}

	return clubs, nil
}

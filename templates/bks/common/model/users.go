package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Mobile   string `gorm:"uniqueIndex;size:20" json:"mobile"`
	Password string `gorm:"size:128" json:"password"`
	Username string `gorm:"size:64" json:"username"`
	Avatar   string `gorm:"size:255" json:"avatar"`
	Source   string `gorm:"size:32" json:"source"`
	Status   int    `gorm:"default:1" json:"status"`
}

func (u *Users) UsersCreate(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *Users) UsersFindById(db *gorm.DB, id uint) error {
	return db.First(u, id).Error
}

func (u *Users) UsersFindByMobile(db *gorm.DB, mobile string) error {
	return db.Where("mobile = ?", mobile).First(u).Error
}

func (u *Users) UsersUpdate(db *gorm.DB) error {
	return db.Save(u).Error
}

func (u *Users) UsersDelete(db *gorm.DB, id uint) error {
	return db.Delete(u, id).Error
}

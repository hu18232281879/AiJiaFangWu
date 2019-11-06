package model

import "github.com/jinzhu/gorm"

func GetAllArea(db *gorm.DB) ([]Area, error) {
	var areas []Area
	err := db.Find(&areas).Error
	return areas, err
}

func InsertUser(db *gorm.DB, user User) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func QueryUser(db *gorm.DB, mobile, hashPassword string) (error,string) {
	var user User
	user.Mobile = mobile
	user.Password_hash = hashPassword
	err:= db.Find(&user).Error
	return err,user.Name
}

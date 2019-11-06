package model

import "github.com/jinzhu/gorm"

func GetAllArea(db *gorm.DB)([]Area,error){
	var areas []Area
	err:=db.Find(&areas).Error
	return areas,err
}
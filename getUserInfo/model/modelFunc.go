package model

import "github.com/jinzhu/gorm"

func GetAllArea(db *gorm.DB)([]Area,error){
	var areas []Area
	err:=db.Find(&areas).Error
	return areas,err
}
func ObtainUserInfo(db *gorm.DB,name string)(User,error){
	var user User
	err:=db.Where("name=?",name).First(&user).Error
	if err!=nil{
		return user,err
	}
	return user,nil
}

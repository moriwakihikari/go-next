package model

import (
	"gorm.io/gorm"
)

type PostEntity struct {
	gorm.Model
	title string
	body  string
}

func GetAll() (datas []PostEntity) {
	result := Db.Find(&datas)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOne(id int) (data PostEntity) {
	result := Db.First(&data, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func (b *PostEntity) Create() {
	result := Db.Create(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *PostEntity) Update() {
	result := Db.Save(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *PostEntity) Delete() {
	result := Db.Delete(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

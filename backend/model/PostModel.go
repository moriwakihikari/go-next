package model

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	id    uint   `json:"id" binding:"required"`
	Title string `json:"title"`
	Body  string `json:"body"`
	// createdAt *time.Time
	// updatedAt *time.Time
	// deletedAt *time.Time
}

func GetAll() (datas []Posts) {
	result := Db.Find(&datas)
	if result.Error != nil {
		panic(result.Error)
	}
	return datas
}

func GetOne(id int) (data Posts) {
	result := Db.First(&data, id)
	log.Printf("%#v", id)

	if result.Error != nil {
		panic(result.Error)
	}
	return data
}

func (p *Posts) Create(c *gin.Context) {
	//postの値挿入される
	// post := Posts{Title: "Jinzhu", Body: "test"}
	log.Printf("%#v", p)
	log.Printf("%#v", c)
	if err := c.BindJSON(&p); err != nil {
		return
	}
	result := Db.Create(&p)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *Posts) Update() {
	result := Db.Save(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *Posts) Delete() {
	result := Db.Delete(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

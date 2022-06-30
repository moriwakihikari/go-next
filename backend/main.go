package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。

	// sqlDB, _ := sql.Open("mysql", "mydb_dsn")
	// gormdb, _ := gorm.Open(mysql.New(mysql.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})

	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	// result := gormdb.Create(&user) // pass pointer of data to Create

	// user.ID             // returns inserted data's primary key
	// result.Error        // returns error
	// result.RowsAffected // returns inserted records count

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// db, err := sql.Open("mysql", "user:password@/dbname")
	// if err != nil {
	// 	panic(err)
	// }
	// // See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)
}

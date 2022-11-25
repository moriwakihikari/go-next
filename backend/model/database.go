package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	// user := os.Getenv("MYSQL_USER")
	// pw := os.Getenv("MYSQL_PASSWORD")
	// db_name := os.Getenv("MYSQL_DATABASE")
	// var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true", "root", "root", "go_next")
	// dialector := mysql.Open(path)
	// var err error
	// if Db, err = gorm.Open(dialector); err != nil {
	// 	connect(dialector, 100)
	// }
	// fmt.Println("db connected!!")
	var err error

	// dsn := "root:root@tcp(127.0.0.1:3306)/go_next?charset=utf8mb4&parseTime=True&loc=Local"
	// Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	Db, err = gorm.Open(mysql.Open("root:root@tcp(localhost)/go_next?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("Could not connect with database!")
	}

	autoMigration()

}

// func connect(dialector gorm.Dialector, count uint) {
// 	var err error
// 	if Db, err = gorm.Open(dialector); err != nil {
// 		if count > 1 {
// 			time.Sleep(time.Second * 2)
// 			count--
// 			fmt.Printf("retry... count:%v\n", count)
// 			connect(dialector, count)
// 			return
// 		}
// 		panic(err.Error())
// 	}
// }
// GetDB is called in models
func GetDB() *gorm.DB {
	return Db
}

// Close is closing db
// func Close() {
// 	if err := Db.Close(); err != nil {
// 		panic(err)
// 	}
// }

func autoMigration() {
	Db.AutoMigrate(&Posts{})
}

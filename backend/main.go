package main

import (
	"database/sql"
	"log"

	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Posts struct {
	gorm.Model
	id    uint
	title string
	body  string
}

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	r.Run() // 0.0.0.0:8080 でサーバーを立てます。

// 	sqlDB, err := sql.Open("mysql",
// 		"user:password@tcp(127.0.0.1:3306)/go_next")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	gormdb, _ := gorm.Open(mysql.New(mysql.Config{
// 		Conn: sqlDB,
// 	}), &gorm.Config{})

// 	var post Posts
// 	// var posts []Posts

// 	result := gormdb.First(&post)
// 	fmt.Print(result)
// 	// result.RowsAffected // returns count of records found
// 	// result.Error        // returns error or nil

// }

func main() {
	connectOnly()

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // 0.0.0.0:8080 でサーバーを立てます。

}

func connectOnly() {
	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_next")
	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	defer db.Close()

	// 実際に接続する
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("データベース接続完了")
	}

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// SQLの実行
	for rows.Next() {
		var post Posts
		err := rows.Scan(&post.id, &post.title, &post.body)

		if err != nil {
			panic(err.Error())
		}
		log.Println(post.id, post.title, post.body)

	}

}

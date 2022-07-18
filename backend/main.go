package main

import (
	"database/sql"
	"go-next/server"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//time型を使う場合DSNに?parseTime=Trueが必要
// type Posts struct {
// 	id        int
// 	title     string
// 	body      string
// 	createdAt time.Time
// 	updatedAt time.Time
// }

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
// }

func main() {
	// connectOnly()

	router := server.GetRouter()
	router.Run()
	// gormDbConnect()

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
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_next?parseTime=True")
	defer db.Close()

	// 実際に接続する
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("データベース接続完了")
	}

	// rows, err := db.Query("SELECT * FROM posts")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// // SQLの実行
	// for rows.Next() {
	// 	var post Posts
	// 	err := rows.Scan(&post.id, &post.title, &post.body, &post.createdAt, &post.updatedAt)

	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	log.Println(post.id, post.title, post.body, post.createdAt, post.updatedAt)

	// }

}

// func gormDbConnect() {
// 	//値の出力方法がわからない

// 	dsn := "root:root@tcp(localhost:3306)/go_next?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		// ここではエラーを返さない
// 		log.Fatal(err)
// 	}
// 	var post Posts
// 	// post := []Posts{}
// 	result := db.Debug().Take(&post)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// log.Println(post.id, post.title, post.body)
// 	str, _ := json.Marshal(result)
// 	fmt.Printf("%s\n", str)
// 	log.Printf("%T\n", result)
// 	log.Printf("%T\n", post)

// }

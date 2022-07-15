package controller

import (
	"fmt"
	"go-next/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowAllPost(c *gin.Context) {
	datas := model.GetAll()
	c.HTML(200, "index.html", gin.H{"datas": datas})
}

func ShowOnePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOne(id)
	c.HTML(200, "show.html", gin.H{"data": data})
}

func ShowCreatePost(c *gin.Context) {
	c.HTML(200, "create.html", gin.H{})
}

// func CreatePost(c *gin.Context) {
// 	title := c.PostForm("title")
// 	body := c.PostForm("body")
// 	data := model.PostEntity{Title: title, Body: body}
// 	data.Create()
// 	c.Redirect(301, "/")
// }

func ShowEditPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOne(id)
	c.HTML(200, "edit.html", gin.H{"data": data})
}

// func EditPost(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.PostForm("id"))
// 	data := model.GetOne(id)
// 	title := c.PostForm("title")
// 	data.Title = title
// 	body := c.PostForm("body")
// 	data.Body = body
// 	data.Update()
// 	c.Redirect(301, "/")
// }

func ShowCheckDeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOne(id)
	c.HTML(200, "delete.html", gin.H{"data": data})
}

func DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	fmt.Println("delete:", id)
	data := model.GetOne(id)
	data.Delete()
	c.Redirect(301, "/")
}

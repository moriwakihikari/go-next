package controller

import (
	"fmt"
	"go-next/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowAllPost(c *gin.Context) {
	datas := model.GetAll()
	c.JSON(200, datas)
	//c.HTML(200, "index.html", gin.H{"datas": datas})

}

func ShowOnePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOne(id)
	c.JSON(200, data)
	//	c.HTML(200, "show.html", gin.H{"data": data})

}

func ShowCreatePost(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func CreatePost(c *gin.Context) {
	// Title := c.PostForm("title")
	// Body := c.PostForm("body")
	data := model.Posts{
		Model: gorm.Model{},
	}
	data.Create(c)
	c.Redirect(301, "/")
}

func ShowEditPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOne(id)
	c.JSON(200, data)
	// c.HTML(200, "edit.html", gin.H{"data": data})
}

func EditPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	data := model.GetOne(id)
	title := c.PostForm("title")
	data.Title = title
	body := c.PostForm("body")
	data.Body = body
	data.Update()
	c.Redirect(301, "/")
}

func ShowCheckDeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := model.GetOne(id)
	c.JSON(200, data)
}

func DeletePost(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	fmt.Println("delete:", id)
	data := model.GetOne(id)
	data.Delete()
	c.Redirect(301, "/")
}

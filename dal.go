package main

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbOps interface {
	init()
	createUser(user *User)
	getUserByName(username string) *User
	createPost(post *Post)
	createComment(comment *Comment)
	deleteCommentById(id uint)
	createCategory(category *Category)
	getCategoryByTitle(title string) *Category
	listPostsByCategory(category *Category) []*Post
	listCommentsByPost(post *Post) []*Comment
}

type DataAccessLayer struct {
	db *gorm.DB
}

func (dal *DataAccessLayer) init() {
	db_file := "/tmp/blog.db"
	os.Remove(db_file)
	var err error
	dal.db, err = gorm.Open(sqlite.Open(db_file), &gorm.Config{})
	if err != nil {
		panic("Database connection failure")
	}

	dal.db.AutoMigrate(&User{}, &Post{}, &Comment{}, &Category{})

	populateDatabase()
}

func (dal *DataAccessLayer) createUser(user *User) {
	dal.db.Create(&user)
}

func (dal *DataAccessLayer) getUserByName(userName string) *User {
	var user User
	dal.db.Find(&user, "user_name = ?", userName)
	return &user
}

func (dal *DataAccessLayer) createPost(post *Post) {
	dal.db.Create(post)
}

func (dal *DataAccessLayer) createComment(comment *Comment) {
	dal.db.Create(comment)
}

func (dal *DataAccessLayer) createCategory(category *Category) {
	dal.db.Create(category)
}

func (dal *DataAccessLayer) getCategoryByTitle(title string) *Category {
	var category Category
	dal.db.Find(&category, "title = ?", title)
	return &category
}

func (dal *DataAccessLayer) deleteCommentById(id uint) {
	dal.db.Delete(&Comment{}, id)
}

func (dal *DataAccessLayer) listPostsByCategory(category *Category) []*Post {
	var posts []*Post
	dal.db.Model(&category).Association("Posts").Find(&posts)
	return posts
}

func (dal *DataAccessLayer) listCommentsByPost(post *Post) []*Comment {
	var comments []*Comment
	dal.db.Find(&comments, "post_id = ?", post.ID)
	return comments
}

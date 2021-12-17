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
	listUsers() []*User
	createPost(post *Post)
	getPostByID(id *uint) *Post
	createComment(comment *Comment)
	getCommentById(id uint) *Comment
	deleteCommentById(id uint)
	createCategory(category *Category)
	getCategoryByTitle(title string) *Category
	listPostsByCategory(category *Category) []*Post
	listPostsByParentPost(post *Post) []*Post
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
	dal.db.Where("username = ?", userName).Find(&user)
	return &user
}

func (dal *DataAccessLayer) listUsers() []*User {
	var users []*User
	dal.db.Find(&users)
	return users
}

func (dal *DataAccessLayer) createPost(post *Post) {
	dal.db.Create(post)
}

func (dal *DataAccessLayer) getPostByID(id *uint) *Post {
	var post Post
	dal.db.First(&post, id)
	return &post
}

func (dal *DataAccessLayer) createComment(comment *Comment) {
	dal.db.Create(comment)
}

func (dal *DataAccessLayer) createCategory(category *Category) {
	dal.db.Create(category)
}

func (dal *DataAccessLayer) getCategoryByTitle(title string) *Category {
	var category Category
	dal.db.Where("title = ?", title).First(&category)
	return &category
}

func (dal *DataAccessLayer) getCommentById(id uint) *Comment {
	var comment Comment
	dal.db.First(&comment, id)
	return &comment
}

func (dal *DataAccessLayer) deleteCommentById(id uint) {
	dal.db.Delete(&Comment{}, id)
}

func (dal *DataAccessLayer) listPostsByCategory(category *Category) []*Post {
	var posts []*Post
	dal.db.Model(&category).Association("Posts").Find(&posts)
	return posts
}

func (dal *DataAccessLayer) listPostsByParentPost(post *Post) []*Post {
	var posts []*Post
	dal.db.Model(&post).Association("Children").Find(&posts)
	return posts
}

func (dal *DataAccessLayer) listCommentsByPost(post *Post) []*Comment {
	var comments []*Comment
	dal.db.Where("post_id = ?", post.ID).Find(&comments)
	return comments
}

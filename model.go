package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string
	PasswordHash string
	Posts        []Post
	Comments     []Comment
}

type Post struct {
	gorm.Model
	Title      string
	Content    string
	UserID     uint
	User       User
	ParentId   *uint
	Children   []Post      `gorm:"foreignkey:ParentId"`
	Categories []*Category `gorm:"many2many:post_categories;"`
	Comments   []*Comment
}

type Comment struct {
	gorm.Model
	Title    string
	Content  string
	PostID   uint
	ParentId *uint
	UserID   uint
	User     User
	children []Comment `gorm:"foreignkey:ParentId"`
}

type Category struct {
	gorm.Model
	Title   string
	Content string
	Posts   []*Post `gorm:"many2many:post_categories;"`
}

package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Posts    []Post
	Comments []Comment
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
	UserID   uint
	User     User
	ParentId *uint
	children []Comment `gorm:"foreignkey:ParentId"`
	PostID   *uint
}

type Category struct {
	gorm.Model
	Title   string
	Content string
	Posts   []*Post `gorm:"many2many:post_categories;"`
}

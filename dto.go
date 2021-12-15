package main

type Credentials struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type NewPost struct {
	Title      string   `binding:"required"`
	Content    string   `binding:"required"`
	Categories []string `binding:"required"`
}

type NewComment struct {
	Title   string `binding:"required"`
	Content string `binding:"required"`
}

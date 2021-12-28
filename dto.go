package main

type Credentials struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

type NewPost struct {
	*UpdatedPost
	Categories []string `binding:"required"`
}

type UpdatedPost struct {
	Title   string `binding:"required"`
	Content string
}

type NewComment struct {
	Title   string `binding:"required"`
	Content string
}

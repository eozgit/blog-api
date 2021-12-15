package main

func populateDatabase() {
	literature := Category{}
	literature.Title = "Literature"
	literature.Content = "Literature"
	app.dal.createCategory(&literature)

	cinema := Category{}
	cinema.Title = "Cinema"
	cinema.Content = "Cinema"
	app.dal.createCategory(&cinema)

	gandalf := User{}
	gandalf.Username = "gandalf"
	gandalf.Password = "youshallnotpass"
	app.dal.createUser(&gandalf)

	frodo := User{}
	frodo.Username = "frodo"
	frodo.Password = "yourelate"
	app.dal.createUser(&frodo)

	post1 := Post{}
	post1.Title = "Hold out your hand, Frodo. It's quite cool."
	post1.Content = "What do you see? Can you see anything?"
	post1.User = gandalf
	// post.Categories = []*Category{literature, cinema}
	app.dal.createPost(&post1)

	categories := app.dal.db.Model(&post1).Association("Categories")
	categories.Append(&literature)
	categories.Append(&cinema)

	comment1 := Comment{}
	comment1.Title = "Nothing. There's nothing."
	comment1.Content = "Wait... there are markings. It's some form of Elvish, I can't read it."
	comment1.PostID = &post1.ID
	comment1.User = frodo
	app.dal.createComment(&comment1)

	comment2 := Comment{}
	comment2.Title = "There are few who can."
	comment2.Content = "The language is the that of Mordor, which I will not utter here."
	comment2.PostID = &post1.ID
	comment2.User = gandalf
	comment2.ParentId = &comment1.ID
	app.dal.createComment(&comment2)

	post2 := Post{}
	post2.Title = "In the common tongue it reads:"
	post2.Content = "One Ring to Rule Them All. One Ring to Find Them. One Ring to Bring Them All and In The Darkness Bind Them."
	post2.User = gandalf
	post2.Categories = []*Category{&cinema}
	app.dal.createPost(&post2)

	// fmt.Printf("comment count %d\n", len(app.dal.listCommentsByPost(&post1)))

	app.dal.deleteCommentById(comment2.ID)

	// fmt.Printf("comment count %d\n", len(app.dal.listCommentsByPost(&post1)))

	// fmt.Printf("literature posts %d\n", len(app.dal.listPostsByCategory(&literature)))

	// fmt.Printf("cinema posts %d\n", len(app.dal.listPostsByCategory(&cinema)))
}

package main

func seedCategories() {
	literature := Category{}
	literature.Title = "Literature"
	literature.Content = "Literature"
	app.dal.createCategory(&literature)

	cinema := Category{}
	cinema.Title = "Cinema"
	cinema.Content = "Cinema"
	app.dal.createCategory(&cinema)
}

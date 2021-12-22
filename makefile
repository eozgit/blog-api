go:
	go run .

test:
	go test -v .

# Docker commands
build:
	docker build --tag eozgit/blog-api .

run:
	docker run --tty --interactive --publish 8080:8080 eozgit/blog-api

# API calls
register:
	clear
	curl --data '{"Username":"gimli","Password":"noonetossesadwarf"}' \
		--header "Content-Type: application/json" --request POST --silent \
		localhost:8080/register | json_pp

post:
	clear
	curl --data '{"Title":"Certainty Of Death? Small Chance Of Success?","Content":"What Are We Waitin'' For?","Categories":["Literature", "Cinema"]}' \
		--header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --silent \
		localhost:8080/post | json_pp

child-post:
	clear
	curl --data '{"Title": "Roaring Fires, Malt Beer","Content":"Ripe Meat Off The Bone!","Categories":["Literature"]}' \
		--header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --silent \
		localhost:8080/post/3 | json_pp

comment:
	clear
	curl --data '{"Title":"I would have gone with you to the end","Content":"Into the very fires of Mordor"}' \
		--header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --silent \
		localhost:8080/post/1/comment | json_pp

child-comment:
	clear
	curl --data '{"Title":"Shall I describe it to you?","Content":"..or would you like me to find you a box?"}' \
		--header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --silent \
		localhost:8080/comment/1 | json_pp

update-post:
	clear
	curl --data '{"Title":"Great!","Content":"Where are we going?"}' \
		--header "Content-Type: application/json" --request PATCH --user gimli:noonetossesadwarf --silent \
		localhost:8080/post/3 | json_pp

delete-comment:
	clear
	curl --header "Content-Type: application/json" --request DELETE --user gimli:noonetossesadwarf --silent \
		localhost:8080/comment/1 | json_pp

list-posts:
	clear
	curl --header "Content-Type: application/json" --request GET --user gimli:noonetossesadwarf --silent \
		localhost:8080/post?category=literature | json_pp

list-comments:
	clear
	curl --header "Content-Type: application/json" --request GET --user gimli:noonetossesadwarf --silent \
		localhost:8080/post/1/comment | json_pp

# Database queries
show-users:
	clear
	sqlite3 -header -column -echo /tmp/blog.db < ./sql/users.sql

show-posts:
	clear
	sqlite3 -header -column -echo /tmp/blog.db < ./sql/posts.sql

show-comments:
	clear
	sqlite3 -header -column -echo /tmp/blog.db < ./sql/comments.sql

show-categories:
	clear
	sqlite3 -header -column -echo /tmp/blog.db < ./sql/categories.sql

show-posts-categories:
	clear
	sqlite3 -header -column -echo /tmp/blog.db < ./sql/posts_categories.sql

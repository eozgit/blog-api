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
aragorn-signup:
	clear
	curl --data '{"Username":"aragorn","Password":"myfriends-youbowtonoone!"}' \
		--header "Content-Type: application/json" --request POST --silent \
		localhost:8080/register | json_pp

aragorn-post-no-stopping:
	clear
	curl --data '{"Title":"Gentlem,","Content":"we do not stop ''til nightfall.","Categories":["Literature", "Cinema"]}' \
		--header "Content-Type: application/json" --request POST --user aragorn:myfriends-youbowtonoone! --silent \
		localhost:8080/post | json_pp

pippin-signup:
	clear
	curl --data '{"Username":"pippin","Password":"itcomesinpints?"}' \
		--header "Content-Type: application/json" --request POST --silent \
		localhost:8080/register | json_pp

pippin-comment-breaky:
	clear
	curl --data '{"Title":"What about breakfast?","Content":""}' \
		--header "Content-Type: application/json" --request POST --user pippin:itcomesinpints? --silent \
		localhost:8080/post/1/comment | json_pp

aragorn-child-post-already:
	clear
	curl --data '{"Title": "You've already had it.","Content":"","Categories":["Literature"]}' \
		--header "Content-Type: application/json" --request POST --user aragorn:myfriends-youbowtonoone! --silent \
		localhost:8080/post/1 | json_pp

aragorn-update-post:
	clear
	curl --data '{"Title":"Gentlemen,","Content":"we do not stop 'til nightfall."}' \
		--header "Content-Type: application/json" --request PATCH --user aragorn:myfriends-youbowtonoone! --silent \
		localhost:8080/post/3 | json_pp

pippin-comment-second:
	clear
	curl --data '{"Title":"We've had one, yes.","Content":"What about second breakfast?"}' \
		--header "Content-Type: application/json" --request POST --user pippin:itcomesinpints? --silent \
		localhost:8080/post/2/comment | json_pp

merry-signup:
	clear
	curl --data '{"Username":"merry","Password":"howmanydidyoueat?"}' \
		--header "Content-Type: application/json" --request POST --silent \
		localhost:8080/register | json_pp

merry-post-i-dont-think-he-knows:
	clear
	curl --data '{"Title":"I don't think he knows about second breakfast, Pip.","Content":"","Categories":["Literature", "Cinema"]}' \
		--header "Content-Type: application/json" --request POST --user merry:howmanydidyoueat? --silent \
		localhost:8080/post | json_pp

pippin-comment-elevenses:
	clear
	curl --data '{"Title":"What about elevenses? Luncheon? Afternoon tea?","Content":"Dinner? Supper? He knows about them, doesn't he?"}' \
		--header "Content-Type: application/json" --request POST --user pippin:itcomesinpints? --silent \
		localhost:8080/post/3/comment | json_pp

merry-comment-wouldnt-count-on-it:
	clear
	curl --data '{"Title":"I wouldn't count on it.","Content":""}' \
		--header "Content-Type: application/json" --request POST --user merry:howmanydidyoueat? --silent \
		localhost:8080/comment/3 | json_pp

delete-comment:
	clear
	curl --header "Content-Type: application/json" --request DELETE --user merry:howmanydidyoueat? --silent \
		localhost:8080/comment/3 | json_pp

list-literature-posts:
	clear
	curl --header "Content-Type: application/json" --request GET --silent \
		localhost:8080/post?category=literature | json_pp

list-cinema-posts:
	clear
	curl --header "Content-Type: application/json" --request GET --silent \
		localhost:8080/post?category=cinema | json_pp

list-comments-by-post:
	clear
	curl --header "Content-Type: application/json" --request GET --silent \
		localhost:8080/post/3/comment | json_pp

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

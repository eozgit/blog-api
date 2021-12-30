curl --data '{"Username":"aragorn","Password":"myfriends-youbowtonoone!"}' \
	--header "Content-Type: application/json" \
	--request POST \
	--silent \
	localhost:8080/register \
	| json_pp

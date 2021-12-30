curl --data '{"Title": "You''ve already had it.","Categories":["Literature"]}' \
    --header "Content-Type: application/json" \
    --request POST \
    --user aragorn:myfriends-youbowtonoone! \
    --silent \
    localhost:8080/post/1 \
    | json_pp
curl --data '{"Title":"Gentlemen,","Content":"we do not stop ''til nightfall."}' \
    --header "Content-Type: application/json" \
    --request PATCH \
    --user aragorn:myfriends-youbowtonoone! \
    --silent \
    localhost:8080/post/1 \
    | json_pp
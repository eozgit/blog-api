curl --data '{"Title":"Gentlem,","Content":"we do not stop ''til nightfall.","Categories":["Literature", "Cinema"]}' \
    --header "Content-Type: application/json" \
    --request POST \
    --user aragorn:myfriends-youbowtonoone! \
    --silent \
    localhost:8080/post \
    | json_pp
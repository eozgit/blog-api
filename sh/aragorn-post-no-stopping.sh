curl --data '{"Title":"Gentlem,","Content":"we do not stop ''til nightfall.","Categories":["Literature"]}' \
    --header "Content-Type: application/json" \
    --request POST \
    --user aragorn:myfriends-youbowtonoone! \
    --silent \
    localhost:8080/post \
    | json_pp
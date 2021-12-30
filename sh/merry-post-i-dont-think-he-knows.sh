curl --data '{"Title":"I don''t think he knows about second breakfast, Pip.","Categories":["Literature", "Cinema"]}' \
    --header "Content-Type: application/json" \
    --request POST \
    --user merry:howmanydidyoueat? \
    --silent \
    localhost:8080/post \
    | json_pp
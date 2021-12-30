curl --header "Content-Type: application/json" \
    --request DELETE \
    --user merry:howmanydidyoueat? \
    --silent \
    localhost:8080/comment/4 \
    | json_pp
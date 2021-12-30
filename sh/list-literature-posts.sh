curl --header "Content-Type: application/json" \
    --request GET \
    --silent \
    localhost:8080/post?category=literature \
    | json_pp
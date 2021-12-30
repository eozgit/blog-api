curl --header "Content-Type: application/json" \
    --request GET \
    --silent \
    localhost:8080/post/3/comment \
    | json_pp
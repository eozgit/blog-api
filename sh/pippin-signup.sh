curl --data '{"Username":"pippin","Password":"itcomesinpints?"}' \
    --header "Content-Type: application/json" \
    --request POST \
    --silent \
    localhost:8080/register \
    | json_pp
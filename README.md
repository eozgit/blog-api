# blog api

blog api loosely based on the following er diagram [source](https://mysql.tutorials24x7.com/blog/guide-to-design-a-database-for-blog-management-in-mysql)

[![er diagram](er.png)](https://mysql.tutorials24x7.com/blog/guide-to-design-a-database-for-blog-management-in-mysql)


#### Run app
```sh
go run .
```


#### Query

```sh
clear && sqlite3 -header -column -echo /tmp/blog.db \
"/* USERS */ select id, username, password, updated_at from users order by updated_at desc; \
/* 

POSTS */ select id, title, content, user_id, parent_id updated_at from posts order by updated_at desc; \
/* 

COMMENTS */ select id, title, content, post_id, parent_id, user_id, deleted_at, updated_at from comments order by updated_at desc; \
/* 

CATEGORIES */ select id, title, content, updated_at from categories order by updated_at desc; \
/* 

POST-CATEGORY MAP */ select * from post_categories;"
```

#### Register
```sh
curl --header "Content-Type: application/json" --request POST --data '{"Username":"gimli","Password":"noonetossesadwarf"}' --write-out "\n" localhost:8080/register
```

#### Create Post
```sh
curl --header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --data '{"Title":"Certainty Of Death? Small Chance Of Success?","Content":"What Are We Waitin'' For?","Categories":["Literature", "Cinema"]}' --write-out "\n" localhost:8080/post
```

#### Run tests
```sh
go test -v .
```

# blog api

blog api loosely based on the following er diagram [source](https://mysql.tutorials24x7.com/blog/guide-to-design-a-database-for-blog-management-in-mysql)

[![er diagram](er.png)](https://mysql.tutorials24x7.com/blog/guide-to-design-a-database-for-blog-management-in-mysql)


#### Run app
```sh
go run .
```

#### Register
```sh
curl --data '{"Username":"gimli","Password":"noonetossesadwarf"}' --header "Content-Type: application/json" --request POST --write-out "\n" localhost:8080/register
```

#### Create Post
```sh
curl --data '{"Title":"Certainty Of Death? Small Chance Of Success?","Content":"What Are We Waitin'' For?","Categories":["Literature", "Cinema"]}' --header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --write-out "\n" localhost:8080/post
```

#### Create Child Post
```sh
curl --data '{"Title": "Roaring Fires, Malt Beer","Content":"Ripe Meat Off The Bone!","Categories":["Literature"]}' --header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --write-out "\n" localhost:8080/post/3
```

#### Create Comment
```sh
curl --data '{"Title":"I would have gone with you to the end","Content":"Into the very fires of Mordor"}' --header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --write-out "\n" localhost:8080/post/1/comment
```

#### Create Child Comment
```sh
curl --data '{"Title":"Shall I describe it to you?","Content":"..or would you like me to find you a box?"}' --header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --write-out "\n" localhost:8080/comment/1
```

#### Run tests
```sh
go test -v .
```


#### Query

```sh
clear && sqlite3 -header -column -echo /tmp/blog.db \
"/* USERS */ select id, username, password, updated_at from users order by updated_at desc; \
/* 

POSTS */ select id, title, content, user_id, parent_id, updated_at from posts order by updated_at desc; \
/* 

COMMENTS */ select id, title, content, post_id, parent_id, user_id, deleted_at, updated_at from comments order by updated_at desc; \
/* 

CATEGORIES */ select id, title, content, updated_at from categories order by updated_at desc; \
/* 

POST-CATEGORY MAP */ select * from post_categories;"
```

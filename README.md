# blog api

blog api loosely based on the following er diagram [source](https://mysql.tutorials24x7.com/blog/guide-to-design-a-database-for-blog-management-in-mysql)

[![er diagram](er.png)](https://mysql.tutorials24x7.com/blog/guide-to-design-a-database-for-blog-management-in-mysql)


#### Run container
```sh
docker run --tty --interactive --publish 8080:8080 eozgit/blog-api
```

#### Register
```sh
curl --data '{"Username":"gimli","Password":"noonetossesadwarf"}' --header "Content-Type: application/json" --request POST --silent localhost:8080/register | json_pp
```

#### Create Post
```sh
curl --data '{"Title":"Certainty Of Death? Small Chance Of Success?","Content":"What Are We Waitin'' For?","Categories":["Literature", "Cinema"]}' --header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --silent localhost:8080/post | json_pp
```

#### Create Child Post
```sh
curl --data '{"Title": "Roaring Fires, Malt Beer","Content":"Ripe Meat Off The Bone!","Categories":["Literature"]}' --header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --silent localhost:8080/post/3 | json_pp
```

#### Create Comment
```sh
curl --data '{"Title":"I would have gone with you to the end","Content":"Into the very fires of Mordor"}' --header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --silent localhost:8080/post/1/comment | json_pp
```

#### Create Child Comment
```sh
curl --data '{"Title":"Shall I describe it to you?","Content":"..or would you like me to find you a box?"}' --header "Content-Type: application/json" --request POST --user gimli:noonetossesadwarf --silent localhost:8080/comment/1 | json_pp
```

#### Update Post
```sh
curl --data '{"Title":"Great!","Content":"Where are we going?"}' --header "Content-Type: application/json" --request PATCH --user gimli:noonetossesadwarf --silent localhost:8080/post/3 | json_pp
```

#### Delete Comment
```sh
curl --header "Content-Type: application/json" --request DELETE --user gimli:noonetossesadwarf --silent localhost:8080/comment/1 | json_pp
```

#### List Posts by Category
```sh
curl --header "Content-Type: application/json" --request GET --user gimli:noonetossesadwarf --silent localhost:8080/post?category=literature | json_pp
```

#### List Comments by Post
```sh
curl --header "Content-Type: application/json" --request GET --user gimli:noonetossesadwarf --silent localhost:8080/post/1/comment | json_pp
```

#### Run code
```sh
go run .
```

#### Run tests
```sh
go test -v .
```

#### Build container
```sh
docker build --tag eozgit/blog-api .
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

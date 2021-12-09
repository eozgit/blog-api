# blog api

blog api loosely based on the following er diagram

![blog er diagram](er.png)

[source](https://mysql.tutorials24x7.com/blog/guide-to-design-a-database-for-blog-management-in-mysql)

#### Query

```sh
clear && sqlite3 -header -column -echo /tmp/blog.db \
"select id, username, password, updated_at from users order by updated_at desc; \
select id, title, content, user_id, parent_id updated_at from posts order by updated_at desc; \
select id, title, content, post_id, parent_id, user_id, deleted_at, updated_at from comments order by updated_at desc; \
select id, title, content, updated_at from categories order by updated_at desc; \
select * from post_categories;"
```

#### Register
```sh
curl --header "Content-Type: application/json" --request POST --data '{"Username":"legolas","Password":"youhavemybow"}' localhost:8080/register
```

#### Run tests
```sh
go test -v .
```

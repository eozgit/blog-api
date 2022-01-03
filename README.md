# blog api

blog api loosely based on the following er diagram [source](https://mysql.tutorials24x7.com/blog/guide-to-design-a-database-for-blog-management-in-mysql)

[![er diagram](er.png)](https://mysql.tutorials24x7.com/blog/guide-to-design-a-database-for-blog-management-in-mysql)


#### Run container
```sh
docker run --tty --interactive --rm --publish 8080:8080 --name blog-api eozgit/blog-api
```

#### Start interactive shell on container
```sh
docker exec --interactive --tty blog-api bash
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

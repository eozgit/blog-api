# blog api

#### Demo
https://youtu.be/puY2l4nB60E


#### Pull image
```sh
docker pull eozgit/blog-api
```

#### Run container
```sh
docker run --rm --tty --interactive --publish 8080:8080 --name blog-api eozgit/blog-api
```

#### Start interactive shell
```sh
docker exec --interactive --tty blog-api bash
```

---

blogging api loosely based on the following er diagram [source](https://mysql.tutorials24x7.com/blog/guide-to-design-a-database-for-blog-management-in-mysql)

[![er diagram](er.png)](https://mysql.tutorials24x7.com/blog/guide-to-design-a-database-for-blog-management-in-mysql)

---

#### Docker repo
https://hub.docker.com/r/eozgit/blog-api

---

#### Build image
```sh
docker build --tag eozgit/blog-api:latest --tag eozgit/blog-api:YYMMDD .
```

#### Push image
```sh
docker image push eozgit/blog-api --all-tags
```

---

#### Run app
```sh
go run .
```

#### Run tests
```sh
go test -v .
```

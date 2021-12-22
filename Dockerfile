FROM golang

EXPOSE 8080

WORKDIR /app

ADD . /app

RUN apt-get update
RUN apt-get install sqlite3
RUN go get .

CMD [ "make", "go" ]

FROM golang

EXPOSE 8080

WORKDIR /app

ADD . /app

RUN echo . /app/sh/set-alias.sh >> /root/.bashrc
RUN apt-get update
RUN apt-get install sqlite3
RUN go get .

CMD [ "go", "run", "." ]

FROM golang

EXPOSE 8080

RUN apt-get --yes update
RUN apt-get --yes install sqlite3

WORKDIR /app

ADD . /app

RUN echo . /app/sh/set-alias.sh >> /root/.bashrc
RUN go get .

CMD [ "go", "run", "." ]

FROM golang

RUN apt install git

EXPOSE 8080

CMD [ "go", "run", "server.go" ]

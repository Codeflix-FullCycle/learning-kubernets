FROM golang

COPY . .

CMD [ "go", "run", "server.go" ]
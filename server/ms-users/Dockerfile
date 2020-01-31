FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 7777

CMD [ "go", "run", "ms-users.go" ]

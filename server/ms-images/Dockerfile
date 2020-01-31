FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8888

CMD [ "go", "run", "ms-images.go" ]

FROM golang:alpine

RUN apk update && \
    apk upgrade && \
    apk add bash git

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . .

ENTRYPOINT ["go", "run", "main.go"]




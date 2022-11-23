FROM golang:1.19-alpine

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /app

COPY go.mod go.sum main.go ./
COPY src ./src/

RUN apk update && apk add --no-cache git
RUN go get ./...

RUN go build -o main .

EXPOSE $PORT

ENTRYPOINT ["./main"]
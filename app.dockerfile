FROM golang:1.19-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN apk upgrade --update && apk --no-cache add git

RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air
RUN go mod download && go mod verify

CMD ["air", "-c", ".air.toml"]
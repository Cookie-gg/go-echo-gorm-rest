# FROM golang:1.19

# RUN apk update && apk add git
# RUN go get github.com/cosmtrek/air@v1.29.0

# WORKDIR /app

# # COPY go.mod go.sum ./
# # RUN go mod download && go mod verify

# CMD ["air", "-c", ".air.toml"]


FROM golang:1.17.7-alpine
WORKDIR /app

RUN apk update && apk add git
RUN go get github.com/cosmtrek/air@v1.29.0


COPY go.mod go.sum ./
RUN go mod download && go mod verify


CMD ["air", "-c", ".air.toml"]
# BUILD
FROM golang:1.16-buster AS build

WORKDIR /app
RUN pwd
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN ls

RUN go build -o server

#DEPLOY
FROM gcr.io/distroless/base-debian10

WORKDIR /usr/src/good-ponds

COPY server .

CMD [ "/usr/src/good-ponds/server" ]
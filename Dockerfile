# BUILD
FROM golang:1.19-buster AS build

WORKDIR /app
COPY . ./
RUN go mod download

RUN go build -o server

#DEPLOY
FROM gcr.io/distroless/base-debian10
WORKDIR /usr/src/good-ponds

COPY --from=build server /usr/src/good-ponds

CMD [ "/usr/src/good-ponds/server" ]
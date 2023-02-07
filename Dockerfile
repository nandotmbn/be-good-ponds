# BUILD
FROM golang:1.19-buster AS build

WORKDIR /app
COPY . ./
RUN go mod download

RUN go build -o /docker-server
RUN ls

#DEPLOY
FROM gcr.io/distroless/base-debian10
WORKDIR /usr/src/good-ponds
RUN pwd
RUN ls

COPY --from=build /docker-server /docker-server

CMD [ "/usr/src/good-ponds/server" ]
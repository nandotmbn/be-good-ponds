# BUILD
FROM golang:1.19-buster AS build

WORKDIR /app
COPY . ./
RUN go mod download

RUN go build -o /docker-server
RUN ls

#DEPLOY
FROM gcr.io/distroless/base-debian10
WORKDIR /

EXPOSE 8080
COPY --from=build /docker-server /docker-server
ENTRYPOINT ["/docker-gs-ping"]
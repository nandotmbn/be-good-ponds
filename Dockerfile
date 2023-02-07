FROM gcr.io/distroless/base-debian10
WORKDIR /usr/src/good-ponds
RUN "go build -o server"
COPY server .
CMD [ "/usr/src/good-ponds/server" ]
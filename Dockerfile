FROM gcr.io/distroless/base-debian10
RUN go build -o server
RUN ls
WORKDIR /usr/src/good-ponds
COPY server .
CMD [ "/usr/src/good-ponds/server" ]
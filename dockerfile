FROM golang:latest
RUN mkdir -p /app/logging/

ADD . /app/
WORKDIR /app/cmd/
RUN go build -o main .
CMD ["/app/cmd/main"]
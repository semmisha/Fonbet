FROM golang:latest
RUN mkdir -p /app/external/logging/
ADD . /app/
WORKDIR /app/cmd
RUN go build -o main .
CMD ["/app/cmd/main"]
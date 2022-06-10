FROM golang:latest
RUN mkdir /app/
RUN mkdir /app/logging/
ADD . /app/
WORKDIR /app/cmd/
RUN go build -o main .
CMD ["/app/cmd/main"]
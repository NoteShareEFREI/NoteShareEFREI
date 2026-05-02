FROM golang:1.26-alpine

WORKDIR /app
ENV GOEXPERIMENT=jsonv2
COPY . .

RUN go mod tidy
RUN go build -o app

EXPOSE 8080

CMD ["./app"]

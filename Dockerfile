FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go get NoteShareEFREI
RUN go build -o app

EXPOSE 8080

CMD ["./app"]

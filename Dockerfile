FROM golang:1.15.8
WORKDIR /go/src/app
COPY . .
RUN go mod download
CMD ["go", "run", "app/main.go"]

EXPOSE 80

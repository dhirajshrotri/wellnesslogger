FROM golang:1.20.5

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app .
CMD ["./app"]
FROM golang:1.21.0

LABEL author="Santio <https://github.com/santiomc>"
WORKDIR /home/container

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /home/container/exporter ./cmd/main.go
EXPOSE 3000

CMD ["/home/container/exporter"]
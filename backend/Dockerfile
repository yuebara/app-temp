FROM golang:1.21

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

EXPOSE 8080

CMD ["air","-c",".air.toml"]
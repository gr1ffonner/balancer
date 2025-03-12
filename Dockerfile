FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go mod tidy && go build -o /bin/app ./cmd

FROM alpine:3.18

COPY --from=builder /bin/app /bin/app

RUN chmod +x /bin/app

CMD ["/bin/app"]
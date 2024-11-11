FROM golang:alpine3.20 as builder

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# RUN go build -o ./bin/bookit ./cmd/bookit/main.go

# FROM alpine:3.20
# WORKDIR /app
# COPY --from=builder /app/bin/bookit /app/
# CMD ["/app/bookit"]

CMD ["air"]

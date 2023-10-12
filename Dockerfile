FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
COPY ./internal/ ./internal/
COPY ./cmd/ ./cmd/

RUN go mod tidy
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./app ./cmd/app/main.go
CMD ["./app"]
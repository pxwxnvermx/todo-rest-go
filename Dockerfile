FROM golang:1.16-alpine as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -o todo-rest ./cmd/main.go

### 

FROM debian:buster-slim

COPY --from=builder /app/todo-rest /app/todo-rest

CMD ["/app/todo-rest"]
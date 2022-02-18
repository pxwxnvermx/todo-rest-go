dev:
	go run ./cmd/main.go

build:
	go build -o todo_rest ./cmd/main.go 

test:
	go test ./...

coverage:
	go test -cover ./...
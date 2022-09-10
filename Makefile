init:
	cp .env.dist .env
	go run ./cmd/main.go

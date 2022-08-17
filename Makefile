init:
	cp .env.dist .env
app-build:
	docker-compose run --rm ws-server go build -o ws-server ./cmd/main.go
app-run:
	docker-compose run --rm ws-server go run ./cmd/main.go

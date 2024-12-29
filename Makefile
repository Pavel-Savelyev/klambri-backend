run:
	go run ./cmd/main.go
deploy-nats-server:
	docker run --rm -p 4222:4222 -p 8222:8222 --name nats-server nats:latest
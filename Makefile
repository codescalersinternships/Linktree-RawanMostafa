include .env

up:
	@echo "Starting container..."
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stopping container..."
	docker-compose down

build:
	go build -o ${BINARY} ./cmd/api/

start: 
	./${BINARY}

restart: build start
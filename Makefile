server-up:
	docker compose -f ./docker-compose.local.yaml up --build --wait

server-down:
	docker compose -f ./docker-compose.local.yaml down > /dev/null 2>&1

server-restart:
	make server-down
	make server-up

api:
	go run cmd/api/main.go
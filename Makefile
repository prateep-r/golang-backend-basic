server-up:
	docker compose -f ./docker-compose.local.yaml up --build --wait

server-down:
	docker compose -f ./docker-compose.local.yaml down > /dev/null 2>&1

init-test:
	go install go.uber.org/mock/mockgen@latest
	go get go.uber.org/mock/mockgen/model

test:
	go test -race -short -coverpkg=training/app/... -coverprofile coverage.out ./...

server-restart:
	make server-down
	make server-up

api:
	go run cmd/api/main.go

producer:
	go run cmd/producer/main.go

consumer:
	go run cmd/consumer/main.go

redis:
	go run cmd/redis/main.go

gen-mock:
	find . -name "mock.go" -type f -delete
	mockgen -destination app/product/mock.go -package=product training/app/product Repository

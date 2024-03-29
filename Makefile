.PHONY: run
run:
	go run ./cmd/todolist

.PHONY: build
build:
	go build ./cmd/todolist

.PHONY: test
test:
	go test ./...

.PHONY: mock
mock:
	mockery --all --keeptree --output=test/mocks

.PHONY: swag
swag:
	swag fmt -g ./cmd/todolist/main.go
	swag init -g ./cmd/todolist/main.go

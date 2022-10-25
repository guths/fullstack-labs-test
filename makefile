migrate-up:
	go run cli.go migrate up

migrate-down:
	go run cli.go migrate down

install:
	go get ./...

prepare: install migrate-up

start:
	go run app/main.go

test:
	go test ./app/tests/...

lint:
	golangci-lint run ./...

challenge-1:
	go run problem-solving/challenge-1/main.go

challenge-2:
	go run problem-solving/challenge-2/main.go

challenge-3:
	go run problem-solving/challenge-3/main.go
install:
	go get ./...

prepare: install

start:
	go run app/main.go

challenge-1:
	go run problem-solving/challenge-1/main.go

challenge-2:
	go run problem-solving/challenge-2/main.go

challenge-3:
	go run problem-solving/challenge-3/main.go

test:
	go test ./app/tests/...

lint:
	golangci-lint run ./...
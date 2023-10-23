build:
	go build -o ./bin/s

run: build
	./bin/s

test:
	go test ./...
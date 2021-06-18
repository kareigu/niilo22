build:
	go build -o bin/main src/main.go

run:
	go run src/main.go

clean:
	go clean

docker:
	docker build -t mxr/niilo-bot:latest .

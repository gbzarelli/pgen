dependencies:
	go get
	go mod tidy

build:
	go build

build-docker:
	docker build -t helpdev/pgen .

run:
	go build
	go run main.go

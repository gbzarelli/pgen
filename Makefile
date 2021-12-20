cmd_path = ./cmd/rest-server

dependencies:
	go get $(cmd_path)
	go mod tidy

build:
	go build $(cmd_path)

build-docker:
	docker build -t helpdev/pgen .

run:
	go build $(cmd_path)
	go run $(cmd_path)

run-docker-compose:
	docker-compose -f .docker-compose/docker-compose.yml up --build

run-docker-stack:
	docker-compose -f .docker-compose/docker-compose-stack.yml up --build

test:
	go test -race ./...

build:
	go build ./cmd/upfluencett/

run: build
	./upfluencett -upfluence-url=${UPFLUENCE_URL} -server-port=${SERVER_PORT}

run-docker: build-docker
	docker run -it --rm --name upfluence-tt-worker upfluence-tt

build-docker:
	docker build -t upfluence-tt . 

run-docker-linux: build-docker-linux
	docker run -it --rm --network=host --name upfluence-tt-worker upfluence-tt

build-docker-linux:
	docker build -t upfluence-tt .

.PHONY: test build run run-docker build-docker run-docker-linux build-docker-linux
-include .env

build:
	@echo "  >  Building package..."
	go build -o cmd/${BIN_FILENAME} ${GO_PACKAGE_NAME}

run:
	@echo "  > Running package..."
	CASSANDRA_USERNAME=${CASSANDRA_USERNAME} \
	CASSANDRA_PASSWORD=${CASSANDRA_PASSWORD} \
	CASSANDRA_HOSTNAME=${CASSANDRA_HOSTNAME} \
	CASSANDRA_KEYSPACE=${CASSANDRA_KEYSPACE} \
	APP_PORT=${APP_PORT} \
	go run ${GO_PACKAGE_NAME}

kill:
	@echo "  >  Killing server..."
	lsof -t -i tcp:8080 | xargs kill

detect-race:
	@echo "  >  Running package in race detection mode..."
	go run -race ${GO_PACKAGE_NAME}

test:
	@echo "  >  Testing package..."
	go test ${GO_PACKAGE_NAME}

fmt:
	@echo "  >  Formating package..."
	go fmt ${GO_PACKAGE_NAME}

clean:
	@echo "  >  Cleaning up build artifacts..."
	go clean

compile:
	@echo "  >  Building package binaries for different platforms..."
	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o cmd/${PROJECT_NAME}-freebsd-386 ${GO_PACKAGE_NAME}
	# MacOS
	GOOS=darwin GOARCH=386 go build -o cmd/${PROJECT_NAME}-darwin-386 ${GO_PACKAGE_NAME}
	# Linux
	GOOS=linux GOARCH=386 go build -o cmd/${PROJECT_NAME}-linux-386 ${GO_PACKAGE_NAME}
	# Windows
	GOOS=windows GOARCH=386 go build -o cmd/${PROJECT_NAME}-windows-386 ${GO_PACKAGE_NAME}
    # 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o cmd/${PROJECT_NAME}-freebsd-amd64 ${GO_PACKAGE_NAME}
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o cmd/${PROJECT_NAME}-darwin-amd64 ${GO_PACKAGE_NAME}
	# Linux
	GOOS=linux GOARCH=amd64 go build -o cmd/${PROJECT_NAME}-linux-amd64 ${GO_PACKAGE_NAME}
	# Windows
	GOOS=windows GOARCH=amd64 go build -o cmd/${PROJECT_NAME}-windows-amd64 ${GO_PACKAGE_NAME}

lint: fmt
	golangci-lint run ./...

visualize:
	go-callvis ${GO_PACKAGE_NAME}

docker-build:
	@echo "  >  Building docker images"
	docker build -t ${DOCKER_IMAGE_NAME} .

docker-run: docker-build
	@echo "  >  Running docker images"
	docker run --env-file .env -p ${APP_PORT}:${APP_PORT} ${DOCKER_IMAGE_NAME}

db-start:
	@echo "  >  Starting up Cassandra"
	docker run --name cassandra -v ${PWD}/cassandra:/var/lib/cassandra -d cassandra -p 9042:9042 --rm

.PHONY: compose
compose:
	@echo "  >  Running containers in docker-compose mode"
	docker-compose up -d

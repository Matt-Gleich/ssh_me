##########
# Building
##########

build-docker-prod:
	docker build -f docker/Dockerfile -t mattgleich/ssh_me:latest .
build-docker-dev:
	docker build -f docker/dev.Dockerfile -t mattgleich/ssh_me:test .
build-docker-dev-lint:
	docker build -f docker/dev.lint.Dockerfile -t mattgleich/ssh_me:lint .
build-go:
	go get -v -t -d ./...
	go build -v .
	rm ssh_me
run-dev:
	python3 ./scripts/reset.py
	SSH_ME_PORT=":2222" go run main.go

#########
# Linting
#########

lint-golangci:
	golangci-lint run
lint-gomod:
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum
lint-goreleaser:
	goreleaser check
lint-hadolint:
	hadolint docker/Dockerfile
	hadolint docker/dev.Dockerfile
	hadolint docker/dev.lint.Dockerfile
lint-in-docker: build-docker-dev-lint
	docker run mattgleich/ssh_me:lint

#########
# Testing
#########

test-go:
	go get -v -t -d ./...
	go test -v ./...
test-in-docker: build-docker-dev
	docker run mattgleich/ssh_me:test

##########
# Grouping
##########

# Testing
local-test: test-go
docker-test: test-in-docker
# Linting
local-lint: lint-golangci lint-goreleaser lint-hadolint lint-gomod
docker-lint: lint-in-docker
# Build
local-build: build-docker-prod build-docker-dev build-docker-dev-lint

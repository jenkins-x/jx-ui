ORG := jenkins-x
BINARY_NAME := jx-ui
LATEST := latest

.PHONY: clean
clean: ## Clean the generated artifacts
	rm -rf build release dist

backend: clean # Builds the go binary
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/linux/$(BINARY_NAME) cmd/app/main.go
backend-darwin: clean # Builds the go binary
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/darwin/$(BINARY_NAME) cmd/app/main.go
	

frontend: # Build the frontend
	npm install --prefix web
	npm run build --prefix web
all: backend frontend

linux: backend

.PHONY: lint
lint: ## Lint the code
	./hack/gofmt.sh
	./hack/linter.sh
build.docker.local: all
	docker build -t ${ORG}/${BINARY_NAME}:${LATEST} .

release: clean linux

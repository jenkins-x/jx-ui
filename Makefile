ORG := jenkins-x
BINARY_NAME := jx-ui
LATEST := latest

clean:
	rm -rf build/

backend: clean # Builds the go binary
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/linux/$(BINARY_NAME) cmd/app/main.go

frontend: # Build the frontend
	npm install --prefix web
	npm run build --prefix web
all: backend frontend

.PHONY: lint
lint: ## Lint the code
	./hack/gofmt.sh
	./hack/linter.sh
build.docker.local: all
	docker build -t ${ORG}/${BINARY_NAME}:${LATEST} .

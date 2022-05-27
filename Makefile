clean:
	rm -rf bin/

backend: clean # Builds the go binary
	go build -o bin/ui cmd/app/main.go

frontend: # Build the frontend
	npm install --prefix web
	npm run build --prefix web
all: backend frontend

.PHONY: lint
lint: ## Lint the code
	./hack/gofmt.sh
	./hack/linter.sh

.PHONY: test
test: ## Runs any tests in the current directory tree
	go test -v -race -cover -parallel 1 ./...

.PHONY: fmt
fmt: ## Formats the code
	go fmt ./...

.PHONY: run-example
run-example: ## Runs the example
	go run ./example/main.go

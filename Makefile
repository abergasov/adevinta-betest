run: ## build app and run parser
	$(info building...)
	go run ./cmd/parser/main.go -path=$(shell pwd)
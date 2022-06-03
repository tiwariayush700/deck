###vendor dependencies
vendor:
	go mod vendor
####run linter
lint:
	golangci-lint run
####run all tests
test:
	go test ./...
####
run:
	go run main/*.go -file=local.json

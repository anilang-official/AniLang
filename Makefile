.PHONY: test
test:
	go test -v ./test
.PHONY: run
run:
	go run main.go
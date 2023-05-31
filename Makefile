.PHONY: test
test:
	go test -v ./test
.PHONY: run
run:
	go run main.go
.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 go build -o bin/macOS/amd/tatakae main.go
	GOOD=darwin GOARCH=arm64 go build -o bin/macOS/arm/tatakae main.go
	GOOS=linux GOARCH=386 go build -o bin/linux/386/tatakae main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd/tatakae main.go
	GOOS=linux GOARCH=arm go build -o bin/linux/arm/tatakae main.go
	GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/tatakae main.go
	GOOS=linux GOARCH=mipsle go build -o bin/linux/mipsle/tatakae main.go
	GOOS=linux GOARCH=ppc64 go build -o bin/linux/ppc64/tatakae main.go
	GOOS=linux GOARCH=ppc64le go build -o bin/linux/ppc64le/tatakae main.go
	GOOS=linux GOARCH=riscv64 go build -o bin/linux/riscv64/tatakae main.go
	GOOS=linux GOARCH=s390x go build -o bin/linux/s390x/tatakae main.go
	GOOS=windows GOARCH=386 go build -o bin/windows/386/tatakae.exe main.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows/amd/tatakae.exe main.go
	GOOS=windows GOARCH=arm go build -o bin/windows/arm/tatakae.exe main.go
	GOOS=windows GOARCH=arm64 go build -o bin/windows/arm64/tatakae.exe main.go
.PHONY: clean
clean:
	rm -rf bin
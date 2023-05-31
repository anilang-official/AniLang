.PHONY: test
test:
	go test -v ./test
.PHONY: run
run:
	go run main.go
.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 go build -o build/macOS/amd64/AniLang/bin/tatakae main.go
	GOOD=darwin GOARCH=arm64 go build -o build/macOS/arm64/AniLang/bin/tatakae main.go
	GOOS=linux GOARCH=386 go build -o build/linux/386/AniLang/bin/tatakae main.go
	GOOS=linux GOARCH=amd64 go build -o build/linux/amd64/AniLang/bin/tatakae main.go
	GOOS=linux GOARCH=arm go build -o build/linux/arm/AniLang/bin/tatakae main.go
	GOOS=linux GOARCH=arm64 go build -o build/linux/arm64/AniLang/bin/tatakae main.go
	GOOS=linux GOARCH=mipsle go build -o build/linux/mipsle/AniLang/bin/tatakae main.go
	GOOS=linux GOARCH=ppc64 go build -o build/linux/ppc64/AniLang/bin/tatakae main.go
	GOOS=linux GOARCH=ppc64le go build -o build/linux/ppc64le/AniLang/bin/tatakae main.go
	GOOS=linux GOARCH=riscv64 go build -o build/linux/riscv64/AniLang/bin/tatakae main.go
	GOOS=linux GOARCH=s390x go build -o build/linux/s390x/AniLang/bin/tatakae main.go
	GOOS=windows GOARCH=386 go build -o build/windows/386/AniLang/bin/tatakae.exe main.go
	GOOS=windows GOARCH=amd64 go build -o build/windows/amd64/AniLang/bin/tatakae.exe main.go
	GOOS=windows GOARCH=arm go build -o build/windows/arm/AniLang/bin/tatakae.exe main.go
	GOOS=windows GOARCH=arm64 go build -o build/windows/arm64/AniLang/bin/tatakae.exe main.go
.PHONY: clean
clean:
	rm -rf bin
.PHONY: macOs-installer
macOS-installer:
	mkdir -p installer/macOs/arm64
	mkdir -p installer/macOs/x64
	pkgbuild --root build/macOS/arm64/AniLang/bin --identifier anilang.arm64.pgk --version 1.0.0 --install-location $(HOME)/AniLang/bin installer/macOs/arm64/tatakae.pkg --scripts scripts/macOs
	pkgbuild --root build/macOS/amd64/AniLang/bin --identifier anilang.amd64.pgk --version 1.0.0 --install-location $(HOME)/AniLang/bin installer/macOs/x64/tatakae.pkg --scripts scripts/macOs
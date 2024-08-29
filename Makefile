
.PHONY: build

build:
	@go build -o build/localbrew cmd/localbrew.go
	@chmod +x build/localbrew

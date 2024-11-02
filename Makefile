build:
	@go build -o bin/homelab cmd/homelab/main.go

run: build
	@./bin/homelab

config:
	@bash config.sh

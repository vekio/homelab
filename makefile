VERSION := $(shell cat version)

# FLAGS
VERSIONFLAG := -X 'github.com/vekio/homelab/pkg/homelab.Version=v$(VERSION)'

install:
	go install -ldflags="${VERSIONFLAG}" ./cmd/homelab
build:
	go build -o ./bin/homelab ./cmd/homelab

default: install

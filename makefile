VERSION := $(shell cat VERSION)

# FLAGS
VERSIONFLAG := -X 'github.com/vekio/homelab/internal/homelab.Version=v$(VERSION)'

install:
	go install -ldflags="${VERSIONFLAG}" ./cmd/homelab

default: install

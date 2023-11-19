traefik-config:
	bash traefik.sh
install:
	cd cli && go install ./cmd/homelab && cd -
get:
	cd cli && go get -u && cd -

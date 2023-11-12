#!/usr/bin/env bash
#
# Replace env varibles in traefik config files.

# Loggers
# -----------------------------------------------------------------------------
info() { printf "%b[info]%b %s\n" '\e[0;32m\033[1m' '\e[0m' "$*" >&2; }
warn() { printf "%b[warn]%b %s\n" '\e[0;33m\033[1m' '\e[0m' "$*" >&2; }
erro() { printf "%b[erro]%b %s\n" '\e[0;31m\033[1m' '\e[0m' "$*" >&2; exit 1; }

# Load .env file
# -----------------------------------------------------------------------------
function dotenv () {
  set -a
  [[ -f .env ]] && . .env
  set +a
}

# Replace env variables
# -----------------------------------------------------------------------------
function traefik-config () {

    cp -r ./traefik/config ./tmp
    cp -r ./traefik/certificates ./tmp
    chmod 600 ./tmp/certificates/acme.json

    # traefik.yml
    sed -i 's/DOMAIN/'"${DOMAIN}"'/g' ./tmp/config/traefik.yml
	sed -i 's/TRAEFIK_CERT_EMAIL/'"${TRAEFIK_CERT_EMAIL}"'/g' ./tmp/config/traefik.yml
	sed -i 's/PROJECT/'"${PROJECT}"'/g' ./tmp/config/traefik.yml

    # # auth.yml
	# sed -i 's/PROJECT/'"${PROJECT}"'/g' ./config/dynamic/auth.yml

    # # pihole.yml
	# sed -i 's/DOMAIN/'"${DOMAIN}"'/g' ./config/dynamic/pihole.yml
	# sed -i 's/PIHOLE/'"${PIHOLE}"'/g' ./config/dynamic/pihole.yml

    # # unraid.yml
	# sed -i 's/DOMAIN/'"${DOMAIN}"'/g' ./config/dynamic/unraid.yml
	# sed -i 's/UNRAID/'"${UNRAID}"'/g' ./config/dynamic/unraid.yml

    # # whitelist.yml
	# sed -i 's/MERCANZAIP/'"${MERCANZAIP}"'/g' ./config/dynamic/whitelist.yml

    # ssh spring "rm ${SPRING_VOL_PATH}/traefik/config/dynamic/*.yml"
    # scp -r ./config spring:"${SPRING_VOL_PATH}"/traefik && \
    # rm -rf ./config && \
    # return 0 || return 1

}

# Main
# -----------------------------------------------------------------------------
function main () {
    info "load .env variables" && dotenv
    traefik-config || erro "traefik-config"
    info "traefik-config done"
}

main "$@"

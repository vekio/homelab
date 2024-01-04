#!/usr/bin/env bash
#
# Send configs files to spring server.

# Loggers
# -----------------------------------------------------------------------------
info() { printf "%b[info]%b %s\n" '\e[0;32m\033[1m' '\e[0m' "$*" >&2; }
warn() { printf "%b[warn]%b %s\n" '\e[0;33m\033[1m' '\e[0m' "$*" >&2; }
err() { printf "%b[error]%b %s\n" '\e[0;31m\033[1m' '\e[0m' "$*" >&2; exit 1; }

# Load .env file
# -----------------------------------------------------------------------------
function dotenv () {
  set -a
  [[ -f .env ]] && . .env-pro
  set +a
}

# Replace env variables
# -----------------------------------------------------------------------------
function upload-configs () {
  ssh spring "mkdir -p ${SPRING_CONFIG} \
    ${SPRING_CONFIG}/authelia \
    ${SPRING_CONFIG}/gitea \
    ${SPRING_CONFIG}/jellyfin \
    ${SPRING_CONFIG}/lldap \
    ${SPRING_CONFIG}/traefik"

  scp -r ${CONFIG}/authelia/config spring:${SPRING_CONFIG}/authelia
  scp -r ${CONFIG}/traefik/config spring:${SPRING_CONFIG}/traefik
}

# Main
# -----------------------------------------------------------------------------
function main () {
    info "load .env variables" && dotenv
    upload-configs || err "upload-configs"
    info "upload-configs done"
}

main "$@"

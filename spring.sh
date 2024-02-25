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
  [[ -f ./services/.env-pro ]] && . ./services/.env-pro
  set +a
}

# Replace env variables
# -----------------------------------------------------------------------------
function upload-configs () {
  ssh spring "mkdir -p ${SPRING_VOLUME} \
    ${SPRING_VOLUME}/authelia \
    ${SPRING_VOLUME}/gitea \
    ${SPRING_VOLUME}/immich \
    ${SPRING_VOLUME}/jellyfin \
    ${SPRING_VOLUME}/lldap \
    ${SPRING_VOLUME}/traefik \
    ${SPRING_VOLUME}/protonmail-bridge"

  # rsync -av --exclude='secrets' ${CONFIG}/authelia spring:${SPRING_VOLUME}/authelia
  # rsync -av --exclude='secrets' ${CONFIG}/gitea spring:${SPRING_VOLUME}/gitea
  # rsync -av --exclude='secrets' ${CONFIG}/immich spring:${SPRING_VOLUME}/immich
  # rsync -av --exclude='secrets' ${CONFIG}/jellyfin spring:${SPRING_VOLUME}/jellyfin
  # rsync -av --exclude='secrets' ${CONFIG}/lldap spring:${SPRING_VOLUME}/lldap
  # rsync -av --exclude='secrets' ${CONFIG}/traefik spring:${SPRING_VOLUME}/traefik
  # scp -r ${CONFIG}/protonmail-bridge/ spring:${SPRING_VOLUME}/protonmail-bridge/
}

# Main
# -----------------------------------------------------------------------------
function main () {
    info "load .env variables" && dotenv
    upload-configs || err "upload-configs"
    info "upload-configs done"
}

main "$@"

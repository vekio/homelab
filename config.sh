#!/usr/bin/env bash

# Loggers
# -----------------------------------------------------------------------------
info() { printf "%b[info]%b %s\n" '\e[0;32m\033[1m' '\e[0m' "$*" >&2; }
warn() { printf "%b[warn]%b %s\n" '\e[0;33m\033[1m' '\e[0m' "$*" >&2; }
err() { printf "%b[error]%b %s\n" '\e[0;31m\033[1m' '\e[0m' "$*" >&2; exit 1; }

# Load .env file
# -----------------------------------------------------------------------------
function dotenv () {
  set -a
  [[ -f .env ]] && . .env
  set +a
}

# Load .env-dev file
# -----------------------------------------------------------------------------
function dotenvdev () {
  set -a
  [[ -f .env-dev ]] && . .env-dev
  set +a
}

# Main
# -----------------------------------------------------------------------------
function main () {
    local inventory_file="pro.ini"

    if [[ "$1" == "--dev" ]]; then
        info "loading .env-dev" && dotenvdev
        inventory_file="dev.ini"
    else
        info "loading .env" && dotenv
    fi

    # Configs
    info "traefik config" && ansible-playbook playbooks/traefik.yml -i "$inventory_file"
    info "adguard config" && ansible-playbook playbooks/adguard.yml -i "$inventory_file"
}

main "$@"

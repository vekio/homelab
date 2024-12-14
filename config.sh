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

# Main
# -----------------------------------------------------------------------------
function main () {
    local inventory_file="pro.ini"
    info "loading .env" && dotenv

    # Configs
    info "traefik config" && ansible-playbook playbooks/traefik.yml -i "$inventory_file"
    # info "adguard config" && ansible-playbook playbooks/adguard.yml -i "$inventory_file"
    # info "status config" && ansible-playbook playbooks/status.yml -i "$inventory_file"
    # info "jellyfin config" && ansible-playbook playbooks/jellyfin.yml -i "$inventory_file"
    info "authelia config" && ansible-playbook playbooks/authelia.yml -i "$inventory_file"
    # info "qbittorrent config" && ansible-playbook playbooks/qbittorrent.yml -i "$inventory_file"
    # info "secrets config" && ansible-playbook playbooks/secrets.yml -i "$inventory_file"
}

main "$@"

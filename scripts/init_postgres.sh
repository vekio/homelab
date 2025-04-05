#!/bin/bash

# Set working directory to script location
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$SCRIPT_DIR/.."

# Import utility functions (env loading, logging, etc.)
source "$SCRIPT_DIR/utils.sh"

main() {
    load_env

    STACKS_ROOT="$ROOT_DIR/docker"

    # Select host
    SERVERS=($(find "$STACKS_ROOT" -mindepth 1 -maxdepth 1 -type d -exec basename {} \; | sort))

    SERVER=$(printf "%s\n" "${SERVERS[@]}" | gum choose --header "Select a server")
    if [[ -z "$SERVER" ]]; then
        log_warn "No server selected. Exiting."
        exit 0
    fi

    # Current context
    CURRENT_CONTEXT=$(docker context show)
    log_info "Current Docker context: $CURRENT_CONTEXT"

    log_info "Switching to docker context: $SERVER"
    docker context use "$SERVER" >/dev/null || {
        log_error "Failed to switch to context: $SERVER"
        exit 1
    }

    # Run playbook
    ansible-playbook "$ROOT_DIR/ansible/playbooks/postgres.yml"

    # Restore previous context
    log_info "Restoring docker context: $CURRENT_CONTEXT"
    docker context use "$CURRENT_CONTEXT" >/dev/null
}

main "$@"

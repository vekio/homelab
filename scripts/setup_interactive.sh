#!/bin/bash

# Set working directory to script location
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$SCRIPT_DIR/.."

# Import utility functions (env loading, logging, etc.)
source "$SCRIPT_DIR/utils.sh"

main() {
    load_env

    PLAYBOOK_DIR="$ROOT_DIR/ansible/playbooks"
    PLAYBOOKS=($(find "$PLAYBOOK_DIR" -maxdepth 1 -type f -name "setup_*.yml" -exec basename {} \;))

    if [[ ${#PLAYBOOKS[@]} -eq 0 ]]; then
        log_error "No setup_* playbooks found in $PLAYBOOK_DIR"
        exit 1
    fi

    # log_info "Select a setup playbook to run:"
    SELECTED=$(printf "%s\n" "${PLAYBOOKS[@]}" | gum choose --header "Playbooks:")

    if [[ -z "$SELECTED" ]]; then
        log_warn "No playbook selected. Exiting."
        exit 0
    fi

    log_info "Running $SELECTED..."
    ansible-playbook "$PLAYBOOK_DIR/$SELECTED"

    log_success "Playbook $SELECTED finished."
}

main "$@"

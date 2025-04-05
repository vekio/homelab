#!/bin/bash

# Set working directory to script location
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Import utility functions (env loading, logging, etc.)
source "$SCRIPT_DIR/utils.sh"

main() {
    log_info "Starting compose files synchronization"

    # Load environment variables from .env file
    load_env

    # Validate required variable
    if [[ -z "$HOMELAB_SHARE" ]]; then
        log_error "HOMELAB_SHARE variable is not defined in the .env file"
        exit 1
    fi

    # Run the Ansible playbook with the loaded environment
    log_info "Running Ansible playbook to sync compose.yml files..."

    ansible-playbook ansible/playbooks/sync_all_compose_files.yml

    log_success "Compose files successfully synchronized."
}

# Run main function
main "$@"

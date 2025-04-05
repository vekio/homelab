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

    # Select stack
    STACK_DIR="$STACKS_ROOT/$SERVER/stacks"
    STACKS=($(find "$STACK_DIR" -mindepth 1 -maxdepth 1 -type d -exec basename {} \; | sort))

    STACK=$(printf "%s\n" "${STACKS[@]}" | gum choose --header "Select a stack for $SERVER:")
    if [[ -z "$STACK" ]]; then
        log_warn "No stack selected. Exiting."
        exit 0
    fi

    # Select action
    TASKS_DIR="$ROOT_DIR/ansible/roles/compose/tasks"
    ACTIONS=($(find "$TASKS_DIR" -maxdepth 1 -type f -name "*.yml" -exec basename {} .yml \; | grep -v main | sort))

    ACTION=$(printf "%s\n" "${ACTIONS[@]}" | gum choose --header "Select an action for stack '$STACK':")
    if [[ -z "$ACTION" ]]; then
        log_warn "No action selected. Exiting."
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
    log_info "🚀 Running '$ACTION' on stack '$STACK' for server '$SERVER'..."

    ansible-playbook "$ROOT_DIR/ansible/playbooks/compose_stack_action.yml" \
        --extra-vars "compose_root=$STACK_DIR/$STACK compose_stack=$STACK compose_action=$ACTION"

    # Restore previous context
    log_info "Restoring docker context: $CURRENT_CONTEXT"
    docker context use "$CURRENT_CONTEXT" >/dev/null

    log_success "Action '$ACTION' completed on stack '$STACK' for '$SERVER'."
}

main "$@"

#!/bin/bash

# Load .env file and export its variables to the environment
load_env() {
    ENV_FILE=".env"

    if [[ ! -f "$ENV_FILE" ]]; then
        log_error ".env file not found at: $ENV_FILE"
        exit 1
    fi

    log_info "Loading environment variables from $ENV_FILE"

    # set -a automatically exports all variables defined after it,
    # allowing .env files to behave like normal shell scripts.
    # This supports quotes, variable expansion, and spacing safely.
    set -a
      . "$ENV_FILE"
    set +a

    return 0
}

# Print an info-level log message
log_info() {
    gum log --level info "🔹 $*"
}

# Print a success message using info level (gum doesn't support 'success' directly)
log_success() {
    gum log --level info "✅ $*"
}

# Print a warning message
log_warn() {
    gum log --level warn "⚠️ $*"
}

# Print an error message
log_error() {
    gum log --level error "❌ $*"
}

# Print a debug message only if DEBUG=true is set in the environment
log_debug() {
    [[ "$DEBUG" == "true" ]] && gum log --level debug "🐞 $*"
}

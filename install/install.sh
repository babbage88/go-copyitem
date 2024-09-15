#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

# Variables
REPO_URL="https://github.com/babbage88/go-copyitem.git"  # Replace with your actual repository URL
BINARY_NAME="cpgo"
INSTALL_DIR="/usr/local/bin"
BASH_COMPLETION_DIR="/etc/bash_completion.d"
ZSH_COMPLETION_DIR="/usr/share/zsh/vendor-completions"
TEMP_DIR=$(mktemp -d)
GO_VERSION_REQUIRED="1.23.0"
GO_TEMP_DIR=$(mktemp -d)
GO_INSTALL_DIR="/usr/local/go"

# Functions

# Check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Install wget if not installed
install_pre_reqs() {
    if ! command_exists wget; then
        echo "wget is not installed. Installing wget..."

        if command_exists apt; then
            sudo apt update
            sudo apt install -y wget git
        elif command_exists yum; then
            sudo yum install -y wget git
        elif command_exists dnf; then
            sudo dnf install -y wget git
        elif command_exists zypper; then
            sudo zypper install -y wget git
        elif command_exists pacman; then
            sudo pacman -Qi wget
            sudo pacman -Qi git 
        else
            echo "Unsupported package manager. Please install wget manually."
            exit 1
        fi
    fi
}

# Get the current Go version
get_go_version() {
    if command_exists go; then
        go version | awk '{print $3}' | sed 's/go//'
    else
        echo "0.0.0"
    fi
}

# Install or update Go if needed
install_go() {
    local current_version
    current_version=$(get_go_version)

    echo "Current Go version: $current_version"
    echo "Required Go version: $GO_VERSION_REQUIRED"

    if [[ "$(printf '%s\n' "$GO_VERSION_REQUIRED" "$current_version" | sort -V | head -n1)" == "$GO_VERSION_REQUIRED" ]]; then
        echo "Go version is up-to-date."
    else
        echo "Installing Go $GO_VERSION_REQUIRED..."
        wget "https://go.dev/dl/go${GO_VERSION_REQUIRED}.linux-amd64.tar.gz" -O "$GO_TEMP_DIR/go.tar.gz"
        sudo rm -rf "$GO_INSTALL_DIR"
        sudo tar -C /usr/local -xzf "$GO_TEMP_DIR/go.tar.gz"

        echo "Updating PATH for all users..."
        sudo tee /etc/profile.d/go.sh <<EOF
export PATH=\$PATH:/usr/local/go/bin
EOF

        # Apply changes to the current shell session
        export PATH=$PATH:/usr/local/go/bin

        echo "Go $GO_VERSION_REQUIRED has been installed and PATH updated."
    fi
}

# Clone repository and build binary
clone_and_build() {
    echo "Cloning repository from $REPO_URL..."
    git clone "$REPO_URL" "$TEMP_DIR"

    echo "Building Go binary..."
    cd "$TEMP_DIR" || exit
    go build -o "$BINARY_NAME"

    echo "Binary built: $TEMP_DIR/$BINARY_NAME"
}

# Get the path of the installed binary
get_installed_binary_path() {
    which "$BINARY_NAME"
}

# Install binary to /usr/local/bin or overwrite the existing one
install_binary() {
    local install_path="$1"
    echo "Installing binary to $install_path..."
    sudo cp "$TEMP_DIR/$BINARY_NAME" "$install_path"

    if [[ $? -eq 0 ]]; then
        echo "Binary successfully installed to $install_path"
    else
        echo "Failed to install the binary"
        exit 1
    fi
}

# Install Bash autocompletion
install_bash_completion() {
    echo "Installing Bash autocompletion..."
    sudo cp "$TEMP_DIR/install/bash_autocomplete" "$BASH_COMPLETION_DIR/$BINARY_NAME"

    if [[ $? -eq 0 ]]; then
        echo "Bash autocompletion installed at $BASH_COMPLETION_DIR/$BINARY_NAME"
    else
        echo "Failed to install Bash autocompletion"
        exit 1
    fi
}

# Install Zsh autocompletion if Zsh is installed
install_zsh_completion() {
    if command_exists zsh; then
        echo "Zsh is installed, installing Zsh autocompletion..."
        sudo cp "$TEMP_DIR/install/zsh_autocomplete" "$ZSH_COMPLETION_DIR/_$BINARY_NAME"

        if [[ $? -eq 0 ]]; then
            echo "Zsh autocompletion installed at $ZSH_COMPLETION_DIR/_$BINARY_NAME"
        else
            echo "Failed to install Zsh autocompletion"
            exit 1
        fi
    else
        echo "Zsh is not installed, skipping Zsh autocompletion installation."
    fi
}

# Clean up temporary files
cleanup() {
    echo "Cleaning up temporary files..."
    rm -rf "$TEMP_DIR" "$GO_TEMP_DIR"
}

# Main installation logic
main() {
    # Install git and wget if needed
    install_pre_reqs

    # Install or update Go if needed
    install_go

    # Clone and build the binary
    clone_and_build

    # Check if binary is already installed
    if command_exists "$BINARY_NAME"; then
        local installed_path
        installed_path=$(get_installed_binary_path)
        echo "Binary is already installed at $installed_path. Overwriting with the new build."
        install_binary "$installed_path"
    else
        echo "Binary is not installed. Installing to $INSTALL_DIR."
        install_binary "$INSTALL_DIR/$BINARY_NAME"
    fi

    # Install Bash and Zsh autocompletion
    install_bash_completion
    install_zsh_completion

    # Clean up
    cleanup
}

# Run the installation process
main

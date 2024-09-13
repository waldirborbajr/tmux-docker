# Define variables for project settings
OUTPUT ?= ./bin
BINARY_NAME ?= tmux-docker         # Name of the output binary (default: myapp)
INSTALL_DIR ?= ~/.local/bin        # Directory to install the binary (default: /usr/local/bin)
TARGET_OS ?= linux                 # Target OS (default: linux)
TARGET_ARCH ?= amd64               # Target architecture (default: amd64)
SOURCE_DIR ?= main.go              # Directory containing main.go or main package

# Go environment
GO = go
GO_BUILD = $(GO) build
GO_INSTALL = install
GO_CLEAN = clean

# Compiler and linker flags
DEV_FLAGS = -gcflags="all=-N -l"                             # Development flags (no optimization, with debug info)
PROD_FLAGS = -ldflags="-s -w" -trimpath -buildmode=pie       # Production flags (small size, high performance)
PROD_GOENV = GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) CGO_ENABLED=0 # Environment for production targeting Linux, disabling CGO

# Build for development
dev: 
	@echo "Building for development..."
	$(GO_BUILD) $(DEV_FLAGS) -o $(OUTPUT)/$(BINARY_NAME) $(SOURCE_DIR)

# Build for production targeting Linux
prod:
	@echo "Building for production targeting Linux..."
	$(PROD_GOENV) $(GO_BUILD) $(PROD_FLAGS) -o $(OUTPUT)/$(BINARY_NAME) $(SOURCE_DIR)

# Install the binary to the system (default: /usr/local/bin)
install: $(BINARY_NAME)
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	$(GO_INSTALL) -m 755 $(OUTPUT)/$(BINARY_NAME) $(INSTALL_DIR)

# Uninstall the binary
uninstall:
	@echo "Removing $(INSTALL_DIR)/$(BINARY_NAME)..."
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	$(GO) $(GO_CLEAN)
	rm -f $(OUTPUT)/$(BINARY_NAME)

# Phony targets that don't represent actual files
.PHONY: dev prod install uninstall clean


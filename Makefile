.PHONY: all build install clean uninstall

BINARY := tmux-docker
INSTALL_DIR=${HOME}/.tmux/plugins/tmux-docker
SRC := main.go
OUT_DIR := ./bin
GO_FLAGS := -ldflags="-s -w"  # Strip debug information, reduce binary size
LDFLAGS := -buildmode=pie     # Position-independent executable for security
GCFLAGS := -trimpath           # Reduce paths in binary for smaller size

# Optimization flags
# LDFLAGS=-s -w -extldflags "-static"

all: devbuild

devbuild:
	@echo "Building $(BINARY) without optimizations..."
	GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/$(BINARY) $(SRC)

prodbuild:
	@echo "Building $(BINARY) with optimizations..."
	GOOS=linux GOARCH=amd64 go build $(GO_FLAGS) $(GCFLAGS) $(LDFLAGS) -o $(OUT_DIR)/$(BINARY) $(SRC)

	# CGO_ENABLED=0 GOOS=linux go build -o $(BINARY_NAME) \
	# 	$(GO_FLAGS) \
	# 	-ldflags '$(LDFLAGS)' \
	# 	-a -installsuffix cgo

prodinstall: prodbuild
	@echo "Installing production version..."
	mkdir -p $(INSTALL_DIR)
	cp $(OUT_DIR)/$(BINARY) $(INSTALL_DIR)/
	cp tmux-docker.tmux $(INSTALL_DIR)/

devinstall: devbuild
	@echo "Installing development version..."
	mkdir -p $(INSTALL_DIR)
	cp $(OUT_DIR)/$(BINARY) $(INSTALL_DIR)/
	cp tmux-docker.tmux $(INSTALL_DIR)/

clean:
	@echo "Cleaning up..."
	rm -rf $(OUT_DIR)/*

uninstall:
	@echo "Uninstalling..."
	rm -rf $(INSTALL_DIR)

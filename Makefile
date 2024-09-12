.PHONY: all build install clean

BINARY_NAME=tmux-docker
INSTALL_DIR=${HOME}/.tmux/plugins/tmux-docker

all: build

build:
	go build -o $(BINARY_NAME)

install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY_NAME) $(INSTALL_DIR)/

clean:
	rm -f $(BINARY_NAME)

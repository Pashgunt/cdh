BINARY_NAME=go-dir
INSTALL_PATH=/usr/local/bin
ZSHRC_PATH=$(HOME)/.zshrc
SCRIPT_PATH=./bash/script.sh
SCRIPT_MARKER=cdh

.PHONY: build install zsh

build:
	go build -o $(BINARY_NAME) ./main.go

install: build
	cp ./$(BINARY_NAME) $(INSTALL_PATH)

zsh: install
	@echo "Configuring Zsh..."
	grep -q '$(SCRIPT_MARKER)' $(ZSHRC_PATH) || \
		(sh -c "echo '\n' >> $(ZSHRC_PATH)" && \
		sh -c "cat $(SCRIPT_PATH) >> $(ZSHRC_PATH)" && \
		source $(ZSHRC_PATH))


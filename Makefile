build-code:
	go build -o go-dir ./main.go && \
	cp ./go-dir /usr/local/bin
build-zsh:
	go build -o go-dir ./main.go && \
	cp ./go-dir /usr/local/bin && \
    sudo sh -c "echo '' >> ~/.zshrc" && \
    sudo sh -c "cat ./bash/script.sh >> ~/.zshrc" && \
	source ~/.zshrc

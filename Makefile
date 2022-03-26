build:
	go build -o bin/tracker github.com/daniilty/boxberry-track/cmd/tracker
install:
	cp bin/tracker /usr/local/bin/boxberry-track
uninstall:
	rm /usr/local/bin/boxberry-track

.PHONY: clean build

all: 
	cd server && GOOS=windows GOARCH=amd64 go build -o ../remote-browser-server.exe 
	cd client && go build -o ../remote-browser-client 

clean:
	rm remote-browser-client remote-browser-server.exe
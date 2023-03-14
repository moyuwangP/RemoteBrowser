.PHONY: clean build release-build

all: 
	cd server && GOOS=windows GOARCH=amd64 go build -o ../remote-browser-server.exe 
	cd client && go build -o ../remote-browser-client 

release-build:
	cd server && GOOS=windows GOARCH=amd64 go build -o ../remote-browser-server.exe 
	cd client && GOOS=windows GOARCH=amd64 go build -o ../remote-browser-client-windows-amd64.exe
	cd client && GOOS=linux GOARCH=amd64 go build -o ../remote-browser-client-linux-amd64
	cd client && GOOS=darwin GOARCH=amd64 go build -o ../remote-browser-client-darwin-amd64
	cd client && GOOS=darwin GOARCH=arm64 go build -o ../remote-browser-client-darwin-arm64

clean:
	rm remote-browser-client remote-browser-server.exe
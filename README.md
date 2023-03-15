# RemoteBrowser
A RemoteApp companion app. Allows you to open web urls in RemoteApp application on your local computer.

## Usage
### On RemoteApp Server
1. Place `remote-browser-server.exe` on your RemoteApp server  
2. Register `remote-browser-server.exe` as your default browser app on your RemoteApp server  
```
    ./remote-browser-server.exe -r
```
Select remote-browser-server.exe as your default browser in Windows Settings if needed

3. On your RemoteApp server, run
```cmd
    ./remote-browser-server.exe -d
```
<br>

### On RemoteApp Client computer
- Place `remote-browser-client` on your client computer
On your local computer, run
``` sh
    ./remote-browser-client [RemoteApp Server IP]
```

## Build
clone and cd to this repository, then
```
    make
```
pick `remote-browser-server.exe` and `remote-browser-client`

## TODO:
- [ ] password encryption
- [x] an easy way to register server app as default browser
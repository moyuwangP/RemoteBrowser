# RemoteBrowser
A RemoteApp companion app. Allows you to open web urls in RemoteApp application on your local computer.

## Usage
Place `remote-browser-server.exe` on your RemoteApp server  
Register `remote-browser-server.exe` as your default browser app on your RemoteApp server  
On your RemoteApp server, run
```cmd
    ./remote-browser-server.exe -d
```
Place `remote-browser-client` on your client computer
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
- [] password encryption
- [] an easy way to register server app as default browser
# v2rayRestfulApi

a restful server generated from v2ray's api .proto files by grpc-gateway 

## usage
### -port: 
port for the restful server to listen on
### -endpoint: 
endpoint of the v2ray grpc api, example: 127.0.0.1:10086

## example
### target
run a restful server on port 80, transform data and then reverse proxy all request into grpc server 127.0.0.1:10086
### *nix
``` Shell
./v2rayRestfulApi -port 80 -endpoint 127.0.0.1:10086
```
### powershell
``` PowerShell
.\v2rayRestfulApi.exe -port 80 -endpoint 127.0.0.1:10086
```

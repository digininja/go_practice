# Break Out

Stuff for trying to break out of Citrix and similar sessions

## Cross compiling

To build for windows i386:

```
GOOS=windows GOARCH=amd64 go build canirun.go -o canirun64.exe
```

And for 64 bit:

```
GOOS=windows GOARCH=amd64 go build  -o canirun64.exe  canirun.go
```

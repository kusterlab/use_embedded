# Installation
```
git submodule update --init --recursive
go get github.com/rakyll/statik
go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
```

# Build it
the Makefile takes a while because it does a lot of cross-compilation. Disable it to make it much faster!
```
make build 
```

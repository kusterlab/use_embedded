# Installation
```
git submodule update --init --recursive
go get github.com/rakyll/statik
go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo
```

Please check that the `GOPATH` is part of your `PATH`. The `GOPATH` is by default located at `~/go`.
If you only want to copy the necessary executables to build this project, do
```bash
cp ~/go/bin/statik /usr/local/bin
cp ~/go/bin/goversioninfo /usr/local/bin
```

# Build it
the Makefile takes a while because it does a lot of cross-compilation. Disable it to make it much faster!
```
make build 
```

# Run it
Go to the `builds` folder and start the executable that suits your system.

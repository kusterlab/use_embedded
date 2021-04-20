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

Github doesn't allow public access to the artifacts, but you can see successful builds [here](https://github.com/kusterlab/use_embedded/actions).

# Run it
Go to the `builds` folder and start the executable that suits your system.

## Server useage
Be aware that either your reverse proxy or this executable has to enable CORS if you want to host this executable for proxy purposes.
e.g. for your NGINX
```
location /proxy_ppc/{
     gzip off;
     add_header Access-Control-Allow-Origin *;
     proxy_pass http://localhost:8080/;
}
```

or you enable it in the [proxy](https://github.com/kusterlab/use_embedded/blob/52b6eb6f8d54c1323df26b79fdf28bc1760a2750/main.go). 

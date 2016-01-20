# How to install

This package includes two simple statically-compiled binary applications, one for Linux on an AMD64 architecture, one for Mac OS X.

You can install them by simply copying the appropriate binary to a folder that's on your PATH.  For example, you could run the following in a terminal:

```
cp linux/civo /usr/bin/
```

If you require it for a different CPU architecture or OS, you can cross-compile the Go source code yourself with (assuming you have a working Go 1.5+ installation):

```
go get github.com/absolutedevops/civo
cd $GOPATH/src/github.com/absolutedevops/civo
GOOS=linux GOARCH=amd64 go build civo.go
cp civo /usr/bin/
```

Replacing the `GOOS` and `GOARCH` with any Go 1.5+ [supported value](https://github.com/golang/go/blob/master/src/go/build/syslist.go).

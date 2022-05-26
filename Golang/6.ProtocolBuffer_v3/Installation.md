## OS: Debian/Ubuntu

Download page: https://github.com/protocolbuffers/protobuf/releases

### Install `protoc` at `$HOME/.local`

- Unzip the file under `$HOME/.local` directory:

`unzip protoc-3.15.8-linux-x86_64.zip -d ~/.local`

- Update environment’s path variable to include
  the path to the protoc executable:

`export PATH="$PATH:$HOME/.local/bin"`

### Install the protocol compiler plugins for Go

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

- Copy `protoc-gen-go` & `protoc-gen-go-grpc` to `$GOROOT/bin`:

```
cd $GOROOT/bin
cp $GOPATH/bin/protoc-gen-go .
cp $GOPATH/bin/protoc-gen-go-grpc .
```

### Verification

- Clone the `grpc-go` repo:

`git clone -b v1.46.0 --depth 1 https://github.com/grpc/grpc-go`

- Change to the quick start example directory:

`cd grpc-go/examples/helloworld/helloworld`

- Remove existing `*pb.go` files:

`rm helloworld_grpc.pb.go helloworld.pb.go`

- Regenerate gRPC code:

```
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
helloworld.proto
```

It should generate `helloworld_grpc.pb.go` & `helloworld.pb.go` files.

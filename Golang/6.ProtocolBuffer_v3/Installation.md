## OS: Debian/Ubuntu

Download page: https://github.com/protocolbuffers/protobuf/releases

### Install ```protoc``` at ```/usr/bin``` path

```
cd ~/
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-linux-x86_64.zip
unzip protoc-3.7.1-linux-x86_64.zip
cd /usr/bin
cp ~/protoc-3.7.1-linux-x86_64/bin/protoc .
```

### Copy files from ```ìnclude``` to ```/usr/include```

```
cd /usr/include/
cp -r ~/protoc-3.7.1-linux-x86_64/include/* .
```

### Install the ```proto``` package

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

### Move ```protoc-gen-go``` to ```$GOROOT/bin```

```
cd $GOROOT/bin
cp $GOPATH/bin/protoc-gen-go .
```

### Verification

- Create ```sample.proto``` file:

```
syntax="proto3";

package main;

message Person {
	string name = 1;
	int32 age = 2;
}
```

- Execute the following command:

```
protoc --go_out=. *.proto
```

It should create a file called ```sample.pb.go```.

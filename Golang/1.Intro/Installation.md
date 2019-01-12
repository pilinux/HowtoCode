## OS: Debian/Ubuntu

Download page: https://golang.org/dl/

### Install ```go``` at ```etc``` directory

```
wget https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz
tar zxvf go1.11.4.linux-amd64.tar.gz -C /etc
```

### Set ```GOROOT```, ```PATH``` and ```GOPATH```  ```Environment_Variables```

- ```GOROOT```: Location of ```go``` tools binary distributions. In the previous step, we have installed ```go``` at ```/etc/go``` directory.

- ```PATH```: ```$GOROOT/bin```

- ```GOPATH```: Location of workspace. For this tutorial, we will create a workspace at ```$HOME/go_workspace/go``` directory.

**For active session**

```
export GOROOT=/etc/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOME/go_workspace/go
```

**Automatically set ```Environment_Variables``` every time physical machine is powered on**

```
export GOROOT=/etc/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=$HOME/go_workspace/go
```

Verify: ```go version```


---
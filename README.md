# WTF is that?

This project is a study case of [Micro toolkit](https://github.com/micro/micro)

# Install dependencies

## Install consul
Download consul [here](https://www.consul.io/downloads.html). To know what is consul, check [here](https://www.consul.io/intro/index.html)

Unzip it at `~/bin`

## Install Go

```
# wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
# tar -C /usr/local/ -xzf go1.8.linux-amd64.tar.gz
```

And add to `/etc/profile` the follow line
```
export PATH=$PATH:/usr/local/go/bin
```

If you’re using zsh, add to your `~/.zshrc`
```
export GOPATH=~/go
export PATH=$PATH:/usr/local/go/bin:~/bin:$GOPATH/bin
```

## Install protobuf
```
$ curl -OL https://github.com/google/protobuf/releases/download/v3.0.0-beta-2/protoc-3.0.0-beta-2-linux-x86_64.zip
$ unzip protoc-3.0.0-beta-2-linux-x86_64.zip -d protoc3
$ sudo mv protoc3/protoc /bin/protoc
```

## Install micro

```
$ go get -u github.com/micro/micro
```

# Run and test

## Run Consul Agent

```
$ consul agent -dev
```

## Run sidecar

```
$ micro sidecar
```

to run with proxy, required to http requests

```
$ micro sidecar --handler=proxy
```

## Run RPC Server

**NOTE**: This service require start micro with NO `--handler=proxy`

First chose your language. This project has an example using Ruby and Go

### Go

run all commands inside `go` dir

> Start RPC Server
```
$ go run service.go
```

> Consume RPC Server (using a client)
```
$ go run client.go
```

### Ruby

run all commands inside `ruby` dir

> Start RPC Server
```
$ ruby rpc_server.rb
```

> Consume RPC Server
```
$ ruby rpc_client.rb
```

## Run HTTP Server

**NOTE**: This service require start micro WITH `--handler=proxy`

First chose your language. This project has an example using Ruby and Go

### Ruby

run all commands inside `ruby` dir

> Start
```
$ ruby http_server.rb
```

> Consume
```
$ ruby http_client.rb
```

# Other helpful tips

## Compile .proto

```
$ cd proto
$ protoc --go_out=. *.proto
```

# Next steps

1. Containerize and run with Docker
2. Create a web server with Go
3. Create a publisher and a subscriber with Go

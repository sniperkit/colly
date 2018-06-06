# gRPCurl
[![Build Status](https://travis-ci.org/fullstorydev/grpcurl.svg?branch=master)](https://travis-ci.org/fullstorydev/grpcurl/branches)

`grpcurl` is a command-line tool that lets you interact with gRPC servers. It's
basically `curl` for gRPC servers.

The main purpose for this tool is to invoke RPC methods on a gRPC server from the
command-line. gRPC servers use a binary encoding on the wire
([protocol buffers](https://developers.google.com/protocol-buffers/), or "protobufs"
for short). So they are basically impossible to interact with using regular `curl`
(and older versions of `curl` that do not support HTTP/2 are of course non-starters).
This program accepts messages using JSON encoding, which is much more friendly for both
humans and scripts.

With this tool you can also browse the schema for gRPC services, either by querying
a server that supports [service reflection](https://github.com/grpc/grpc/blob/master/src/proto/grpc/reflection/v1alpha/reflection.proto)
or by loading in "protoset" files (files that contain encoded file
[descriptor protos](https://github.com/google/protobuf/blob/master/src/google/protobuf/descriptor.proto)).
In fact, the way the tool transforms JSON request data into a binary encoded protobuf
is using that very same schema. So, if the server you interact with does not support
reflection, you will need to build "protoset" files that `grpcurl` can use.

This repo also provides a library package, `github.com/fullstorydev/grpcurl`, that has
functions for simplifying the construction of other command-line tools that dynamically
invoke gRPC endpoints. This code is a great example of how to use the various packages of
the [protoreflect](https://godoc.org/github.com/jhump/protoreflect) library, and shows
off what they can do.

## Features
`grpcurl` supports all kinds of RPC methods, including streaming methods. You can even
operate bi-directional streaming methods interactively by running `grpcurl` from an
interactive terminal and using stdin as the request body!

`grpcurl` supports both plain-text and TLS servers and has numerous options for TLS
configuration. It also supports mutual TLS, where the client is required to present a
client certificate.

As mentioned above, `grpcurl` works seamlessly if the server supports the reflection
service. If not, you can supply the `.proto` source files or you can supply protoset
files (containing compiled descriptors, produced by `protoc`) to `grpcurl`.

## Installation
You can use the `go` tool to install `grpcurl`:
```shell
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

This installs the command into the `bin` sub-folder of wherever your `$GOPATH`
environment variable points. If this directory is already in your `$PATH`, then
you should be good to go.

If you have already pulled down this repo to a location that is not in your
`$GOPATH` and want to build from the sources, you can `cd` into the repo and then
run `make install`.

If you encounter compile errors, you could have out-dated versions of `grpcurl`'s
dependencies. You can update the dependencies by running `make updatedeps`.

## Example Usage
Invoking an RPC on a trusted server (e.g. TLS without self-signed key or custom CA)
that requires no client certs and supports service reflection is the simplest thing to
do with `grpcurl`. This minimal invocation sends an empty request body:
```shell
grpcurl grpc.server.com:443 my.custom.server.Service/Method
```

To list all services exposed by a server, use the "list" verb. When using `.proto` source
or protoset files instead of server reflection, this lists all services defined in the
source or protoset files.
```shell
# Server supports reflection
grpcurl localhost:8787 list

# Using compiled protoset files
grpcurl -protoset my-protos.bin list

# Using proto sources
grpcurl -import-path ../protos -proto my-stuff.proto list
```

The "list" verb also lets you see all methods in a particular service:
```shell
grpcurl localhost:8787 list my.custom.server.Service
```

The "describe" verb will print the type of any symbol that the server knows about
or that is found in a given protoset file and also print the full descriptor for the
symbol, in JSON.
```shell
# Server supports reflection
grpcurl localhost:8787 describe my.custom.server.Service.MethodOne

# Using compiled protoset files
grpcurl -protoset my-protos.bin describe my.custom.server.Service.MethodOne

# Using proto sources
grpcurl -import-path ../protos -proto my-stuff.proto describe my.custom.server.Service.MethodOne
```

The usage doc for the tool explains the numerous options:
```shell
grpcurl -help
```

## Proto Source Files
To use `grpcurl` on servers that do not support reflection, you can use `.proto` source
files.

In addition to using `-proto` flags to point `grpcurl` at the relevant proto source file(s),
you may also need to supply `-import-path` flags to tell `grpcurl` the folders from which
dependencies can be imported.

Just like when compiling with `protoc`, you do *not* need to provide an import path for the
location of the standard protos included with `protoc` (which contain various "well-known
types" with a package definition of `google.protobuf`). These files are "known" by `grpcurl`
as a snapshot of their descriptors is built into the `grpcurl` binary.

## Protoset Files
You can also use compiled protoset files with `grpcurl`. If you are scripting `grpcurl` and
need to re-use the same proto sources for many invocations, you will see better performance
by using protoset files (since it skips the parsing and compilation steps with each
invocation).

Protoset files contain binary encoded `google.protobuf.FileDescriptorSet` protos. To create
a protoset file, invoke `protoc` with the `*.proto` files that describe the service:
```shell
protoc --proto_path=. \
    --descriptor_set_out=myservice.protoset \
    --include_imports \
    my/custom/server/service.proto
```

The `--descriptor_set_out` argument is what tells `protoc` to produce a protoset,
and the `--include_imports` argument is necessary for the protoset to contain
everything that `grpcurl` needs to process and understand the schema.
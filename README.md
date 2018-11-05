## IPFS Middleware

A _YOLO_ middleware add-on for [go-ipfs](https://github.com/ipfs/go-ipfs).

### Usage

- install the package: `go get -u github.com/wzulfikar/ipfs-middleware`
- get the go-ipfs repo: `go get -u -d github.com/ipfs/go-ipfs`
- go inside the repo: `cd $GOPATH/src/github.com/ipfs/go-ipfs`
- add the code for middleware handler inside [`core/corehttp/gateway_handler.go:L91`](https://github.com/ipfs/go-ipfs/blob/3f7668bdca719b3cbeac741fd2c5f454b5d14c15/core/corehttp/gateway_handler.go#L91). something like this:

  ```diff
  // gateway_handler.go

  import (
      // import ipfs-middleware
      middleware "github.com/wzulfikar/ipfs-middleware"
  )

  defer func() {
  	if r := recover(); r != nil {
  	  log.Error("A panic occurred in the gateway handler!")
  	  log.Error(r)
  	  debug.PrintStack()
  	}
  }()
  
  +// pass gateway request to middleware
  +if ok := middleware.Handle(w, r); !ok {
  +  return
  +}

  if i.config.Writable {
  	switch r.Method {
  	  case "POST":
  	    i.postHandler(ctx, w, r)
  	    return
     case "PUT":
  	    i.putHandler(w, r)
  	    return
     case "DELETE":
  	    i.deleteHandler(w, r)
  	    return
    }
  }
  ```
- install go-ipfs: `make install`
- build the binaries: `make build`
- run the newly build ipfs binary: `cmd/ipfs/ipfs daemon`

### Verifying the middleware

Once the ipfs daemon is running, you can check if the middleware is working correctly by sending a get request to a gateway path.

First, let's create a file and add it to our ipfs node.

`echo Hi! > hi.txt && ./cmd/ipfs/ipfs add hi.txt && rm hi.txt`

It will give you this hash:

`QmV45zTXVeNkwtKPG37Rt8KY14fhxWFU56fUGzxw1ixJ6r`

Originally, you can access the file from your browser directly, or using curl:

`curl localhost:8080/ipfs/QmV45zTXVeNkwtKPG37Rt8KY14fhxWFU56fUGzxw1ixJ6r`

However, since the currently-running daemon is the one with built with added middleware, the curl request will return `unauthorized`. This is because the ipfs-middleware that we've added has included some example middlewares (see: `./handler.go`). One of the example middleware is the auth middleware (./middlewares/auth/handler.go), that will drop the request if the request doesn't have `Authorization` header.

Now, let's add an Authorization header and it'll greet you with "Hi!":

`curl -H "Authorization: bearer yada-yada" localhost:8080/ipfs/QmV45zTXVeNkwtKPG37Rt8KY14fhxWFU56fUGzxw1ixJ6r`

_That's it :)_

You've added a code that gives middleware-like functionality into ipfs binary.

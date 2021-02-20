# go-hotreload

Simple Go app executed in a container with hot reload and dynamic import of a local package.

## Description

It is a good practice the run Go application in a Docker container, and use hot reload and local packages for developpment is important.

This app allows to:

- Start a Go application in a Docker container.
- Use hot reload to restart the application each time a file is saved in `myapp`.
- Import a dynamicly updatable local package from `mypackage`.

## How to use

## Start

Get this repository:

```
git clone git@github.com:phoax/go-hotreload.git
```

Build images and start app:

```
cd myapp/
make dev
```

Or force build images and start app:

```
cd myapp/
make dev-build
```

The terminal should display:

```
Call Package: Hello World from Mypackage
```

## Test

### Update mypackage

Update `mypackage/mytest/mytest.go ` by changing response message.

Example:

replace `message := ...` by

```
message := "Hello World from Mypackage updated"
```

### Update myapp

Update `myapp/main.go` by changing response message.

Example:

replace `message := ...` by

```
message := "Call updated Package:"
```

When `myapp/main.go` is saved, the app should rebuild, and the terminal should display:

```
Call updated Package: Hello World from Mypackage updated
```

## How it works

### Dynamic local package

Add a new local package only to `go.dev.mod`

```
go mod edit -replace github.com/phoax/go-hotreload/mypackage=../mypackage ./go.dev.mod
```

Then in docker-compose, the `go.mod` will be overwritten with `../myapp/go.dev.mod:/go/src/myapp/go.mod`:

```
[...]
    volumes:
      [...]
      - ../myapp/go.dev.mod:/go/src/myapp/go.mod
```

### Hot reload

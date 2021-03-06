# go-hotreload

Simple Go app executed in a container with hot reload and dynamic import of a local package.

## Description

It is a good practice the run Go application in a Docker container, and use hot reload and local packages for developpment is important.

This app allows to:

- Start a Go application in a Docker container.
- Use hot reload to restart the application each time a file is saved in `myapp`.
- Import a dynamicly updatable local package from `mypackage`.

## How to use

The application can be executed with the released package or with the local package.

### Start with the released package

The application can be executed with the released package on GitHub to preserve compatibility with official packages.

Build image and start app:

```
cd myapp/
make dev
```

Or force build image and start app:

```
cd myapp/
make dev-build
```

The terminal should display:

```
Call Package: Hello World from Mypackage
```

### Start with the local package

The application can be executed with the local package to allow dynamic updates on the package.

Build image and start app:

```
cd myapp/
make dev-local
```

Or force build image and start app:

```
cd myapp/
make dev-local-build
```

The terminal should display:

```
Call Package: Hello World from Mypackage
```

### Test

#### Update mypackage

Update `mypackage/mytest/mytest.go ` by changing response message.

Example:

replace `message := ...` by

```
message := "Hello World from Mypackage updated"
```

#### Update myapp

Update `myapp/main.go` by changing response message.

Example:

replace `message := ...` by

```
message := "Call updated Package:"
```

When `myapp/main.go` is saved, the app should rebuild.

If the app is started with `make dev` or `make dev-build`, the package is not updated, the terminal should display:

```
Call updated Package: Hello World from Mypackage
```

If the app is started with `make dev-local` or `make dev-local-build`, the package is dynamically updated, the terminal should display:

```
Call updated Package: Hello World from Mypackage updated
```

## How it works

### Dynamic local package

To use a local package, specify the `./go.dev.mod` file in the `go mod edit -replace` command.

Example:

```
go mod edit -replace github.com/phoax/go-hotreload/mypackage=../mypackage ./go.dev.mod
```

Then, when started the `go.mod` app will be overwritten by `go.dev.mod` thanks to `../myapp/go.dev.mod:/go/src/myapp/go.mod`

docker-compose.yml:

```
[...]
    volumes:
      [...]
      - ../myapp/go.dev.mod:/go/src/myapp/go.mod
```

### Hot reload

Hot reload is simply enabled by compiling and starting the app with `go get github.com/githubnemo/CompileDaemon` package.

Dockerfile:

```
[...]

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command="./main"
```

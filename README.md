# go-hotreload

Simple Go app executed in a container with hot reload and local packages.

## Description

It is a good practice the run Go application in a Docker container

This app allows to:

- Start a Go application in a Docker container.
- Use hot reload to restart the application each time a new file is saved in Myapp.
- Import an updatable local package from Mypackage.

## Start

Build images and start service

```
cd myapp/
make dev
```

Force build images and start service

```
cd myapp/
make dev-build
```

## Mod

In Mypackage

```
go mod init github.com/phoax/go-hotreload/mypackage
```

In Myapp

```
go mod init github.com/phoax/go-hotreload/myapp
go mod edit -replace github.com/phoax/go-hotreload/mypackage=../mypackage
```

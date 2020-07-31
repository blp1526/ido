[![Build Status](https://travis-ci.org/blp1526/ido.svg?branch=master)](https://travis-ci.org/blp1526/ido)
[![Go Report Card](https://goreportcard.com/badge/github.com/blp1526/ido)](https://goreportcard.com/report/github.com/blp1526/ido)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/blp1526/ido)](https://pkg.go.dev/github.com/blp1526/ido)
[![GolangCI](https://golangci.com/badges/github.com/blp1526/ido.svg)](https://golangci.com/r/github.com/blp1526/ido)

# ido

A Toy Container

## Required commands

* chroot
* unshare
* mount
* umount
* tar
* docker
  * only `docker create`, `docker export` and `docker rm`

## Build

```
$ make
```

## Usage

```
$ sudo ./bin/ido run busybox:latest /bin/sh
```

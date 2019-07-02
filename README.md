# ido

A Toy Container

## Required commands

* chroot
* unshare
* tar
* docker
  * only `docker create`, `docker export` and `docker rm`

## Build

```
$ make
```

## Usage

```
$ sudo ido run busybox:latest /bin/sh
```

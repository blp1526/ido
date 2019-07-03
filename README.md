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

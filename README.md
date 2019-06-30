# ido

## Required commands

* chroot
* unshare
* tar
* docker
  * only `docker create` and `docker export`

## Build

```
$ make
```

## Usage

```
$ sudo ido create busybox:latest
```

```
$ sudo ido run /bin/sh
```

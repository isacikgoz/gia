![](https://img.shields.io/travis/com/isacikgoz/gia.svg?style=flat) 

# gia

`gia` is a tool for interactive hunk staging. 

### Preview

<p align="center">
   <img src="https://github.com/isacikgoz/gia/blob/master/images/preview.png" alt="screenshot"/>
</p>

### Install

```shell
go get github.com/isacikgoz/gia
```

### Usage

```shell
usage: gia [<file>] [<args>]

Flags:
	-h, --help     Show help.
	-v, --version  Show application version.

Args:
  --
    Default mode, used on unstaged files.

  --no-index
    Used on untracked files.

  --cached
    Used on staged files.
```

### License
[BSD-3-Clause](/LICENSE)
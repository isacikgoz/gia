![](https://img.shields.io/travis/com/isacikgoz/gia.svg?style=flat) 

# gia

`gia` is a tool for interactive hunk staging. 

### Preview

<p align="center">
   <img src="https://user-images.githubusercontent.com/2153367/57180688-d5ef7a00-6e93-11e9-819f-315db3bf09ff.gif" alt="screencast"/>
</p>

### Install

```shell
go get github.com/isacikgoz/gia
```

### Program Arguments

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

### Controls

- ↓, j: Cursor down
- ↑, k: Cursor up
- n: Next hunk
- N: Previous hunk
- g: Go to top
- G: Go to bottom
- space: Stage/Unstage
- q: Quit
- c: Open controls

### Adding To Your Own App

See `main.go` for usage. Also, feel free to open PR's.

### License
[BSD-3-Clause](/LICENSE)
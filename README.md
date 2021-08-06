# Example of Go-IPFS As A Library

This repo is a spin-off of the [official tutorial sample of `go-ipfs` (from Go-IPFS v0.7.0)](https://github.com/ipfs/go-ipfs/tree/v0.7.0/docs/examples/go-ipfs-as-a-library).

It aims to be the working example for various Golang version.

- [x] Go 1.14
- [x] Go 1.15
- [x] Go 1.16

## How To Run

```shellsession
$ cd /path/to/cloned/go-ipfs-as-a-library
$ go mod download
...(** snip **)...
$ go run main.go
...(** snip **)...
All done! You just finalized your first tutorial on how to use go-ipfs as a library
```

## How To Test

```shellsession
$ go test .
ok      github.com/ipfs/go-ipfs/examples/go-ipfs-as-a-library   2.362s
```

```bash
# Runs the tests on Go v1.14, 1.15, 1.16 over Alpine Linux
$ docker-compose up
...(** snip **)...
v1_14_1  | ok  	github.com/ipfs/go-ipfs/examples/go-ipfs-as-a-library	2.973s
v1_14_1 exited with code 0
v1_15_1  | ok  	github.com/ipfs/go-ipfs/examples/go-ipfs-as-a-library	2.205s
v1_15_1 exited with code 0
v1_16_1  | ok  	github.com/ipfs/go-ipfs/examples/go-ipfs-as-a-library	1.696s
v1_16_1 exited with code 0
```

## About the tutorial

For the descriptions about the tutorial, see: [REAMDE2.md](README2.md)

## License

- Original Authors:
  - [go-ipfs/commits/v0.7.0/docs/examples/go-ipfs-as-a-library](https://github.com/ipfs/go-ipfs/commits/v0.7.0/docs/examples/go-ipfs-as-a-library)
- [MIT](LICENSE-MIT)/[Apache-2.0](LICENSE-APACHE), [Dual License](LICENSE)
  - See: [Issue #6302](https://github.com/ipfs/go-ipfs/issues/6302) | go-ipfs | ipfs @ GitHub

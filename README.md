# Example of Go-IPFS As A Library

This repo is a spin-off of the [official tutorial sample of `go-ipfs` (from Go-IPFS v0.7.0)](https://github.com/ipfs/go-ipfs/tree/v0.7.0/docs/examples/go-ipfs-as-a-library).

It aims to be the working example for various Golang version.

- Example: [main.go](./main.go)
- Tests of the example: [main_test.go](./main_test.go)
- Results of weekly testing
  - Go 1.14 [![go1_14](https://github.com/KEINOS/go-ipfs-as-a-library/actions/workflows/runGo1_14.yml/badge.svg)](https://github.com/KEINOS/go-ipfs-as-a-library/actions/workflows/runGo1_14.yml)
  - Go 1.15
[![go1_15](https://github.com/KEINOS/go-ipfs-as-a-library/actions/workflows/runGo1_15.yml/badge.svg)](https://github.com/KEINOS/go-ipfs-as-a-library/actions/workflows/runGo1_15.yml)
  - Go 1.16 [![go1_16](https://github.com/KEINOS/go-ipfs-as-a-library/actions/workflows/runGo1_16.yml/badge.svg)](https://github.com/KEINOS/go-ipfs-as-a-library/actions/workflows/runGo1_16.yml)

## How To Run

```shellsession
$ git clone https://github.com/KEINOS/go-ipfs-as-a-library.git
...(** snip **)...

$ cd ./go-ipfs-as-a-library

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

If you have Docker installed, you can run the test with `docker-compose`.

```bash
# It runs the tests on Go v1.14, 1.15, 1.16 over Alpine Linux
$ docker-compose up
...(** snip **)...
v1_14_1  | ok  	github.com/ipfs/go-ipfs/examples/go-ipfs-as-a-library	2.973s
v1_14_1 exited with code 0
v1_15_1  | ok  	github.com/ipfs/go-ipfs/examples/go-ipfs-as-a-library	2.205s
v1_15_1 exited with code 0
v1_16_1  | ok  	github.com/ipfs/go-ipfs/examples/go-ipfs-as-a-library	1.696s
v1_16_1 exited with code 0
```

## References for Go-IPFS

- [Working with Go](https://docs.ipfs.io/reference/go/api/#working-with-go) @ docs.ipfs.io
- [Use go-ipfs as a library to spawn a node and add a file](https://github.com/ipfs/go-ipfs/blob/master/docs/examples/go-ipfs-as-a-library/README.md) | Docs | go-ipfs @ GitHub

## License

- Original Authors:
  - [go-ipfs/commits/v0.7.0/docs/examples/go-ipfs-as-a-library](https://github.com/ipfs/go-ipfs/commits/v0.7.0/docs/examples/go-ipfs-as-a-library)
- [MIT](LICENSE-MIT)/[Apache-2.0](LICENSE-APACHE), [Dual License](LICENSE)
  - See: [Issue #6302](https://github.com/ipfs/go-ipfs/issues/6302) | go-ipfs | ipfs @ GitHub

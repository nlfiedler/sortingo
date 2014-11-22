# README

## Sorting for Go

This is the source distribution of a sorting package for the [Go](http://golang.org) programming language, named sortingo, which contains various sorting algorithm implementations.

## Installation

Install [Git](http://git-scm.com) in order to fetch the other dependencies.

Run the `go` tool like so:

```
go get -t github.com/nlfiedler/sortingo
```

## Benchmarks

To run the benchmarks, first build the one or both benchmarking "commands", like so:

```
$ cd $GOPATH/src/github.com/nlfiedler/sortingo
$ go install ...sortingo/cmd/sortbench
$ go install ...sortingo/cmd/sortmbench
$ $GOPATH/bin/sortbench
...
$ $GOPATH/bin/sortmbench
...
```

## License

The sortingo project is licensed under the [New BSD](http://opensource.org/licenses/BSD-3-Clause) license.

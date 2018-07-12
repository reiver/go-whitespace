# go-whitespace

A small Golang library for dealing with whitespace.

Note that this Go library accounts for all 26 UNICODE whitespace characters (and not just the standard ASCII ones),
as well as accounts for all 7 UNICODE mandatory break characters.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-whitespace

[![GoDoc](https://godoc.org/github.com/reiver/go-whitespace?status.svg)](https://godoc.org/github.com/reiver/go-whitespace)


Usage
-----

To use this Golang library, use with something like:
```go
import "github.com/reiver/go-whitespace"

// ...


if whitespace.IsWhitespace(r) {
	//@TODO: Do something.
}
```

```go
import "github.com/reiver/go-whitespace"

// ...

if whitespace.IsMandatoryBreak(r) {
	//@TODO: Do something.
}

```

Go Whitespace
=============

A small Golang library for dealing with whitespace.

Note that this Go library accounts for all 26 UNICODE whitespace characters (and not just the standard ASCII ones),
as well as accounts for all 7 UNICODE mandatory break characters.


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

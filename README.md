Go Whitespace
=============

A small Golang library for dealing with whitespace.

Note that this Go library accounts for all 26 UNICODE whitespace characters (and not just the standard ASCII ones).


Usage
-----

To use this Golang library, use with something like:
```
package main

import "github.com/reiver/go-whitespace"

func main() {

	// ...

	if whitespace.IsWhitespace(r) {
		// Do something
	}

	// ...
}
```

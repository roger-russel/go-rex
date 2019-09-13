# Go Rex

[![CircleCI](https://circleci.com/gh/roger-russel/go-rex/tree/master.svg?style=shield)](https://circleci.com/gh/roger-russel/go-rex/tree/master)
[![codecov](https://codecov.io/gh/roger-russel/go-rex/branch/master/graph/badge.svg)](https://codecov.io/gh/roger-russel/go-rex)
[![Software License](https://img.shields.io/github/license/roger-russel/go-rex)](LICENSE.md)

Go Rex is a bind to use Regex.h instead of the Golang native regex parser.

Native Go Regex is quit slow, it looses to almost all other languages, why? Because the other languages are using Regex.h lib into it and Golang is not.

The problem is not the Golang, it just that the Golang native library needs be fixed to be more performatic, well until they fix it, we can at last use the regex.h too.

The regex.h use GNU ERE, its mean that character classes are write into this way.
eg: Numeric class is not ```[\d]``` but ```[[:digit:]]```.

Helpful there is a list on wikipedia with all list [here](https://en.wikipedia.org/wiki/Regular_expression#Character_classes]).

## Installing

```sh
dep ensure -add github.com/roger-russel/go-rex
```

## How Use

All examples can be found [here](./docs/examples).

```go
package main

import (
	"fmt"

	"github.com/roger-russel/go-rex/gorex"
)

func main() {

	pattern := "^[a-zA-Z0-9_.+-]+@[a-zA-Z-]+\\.[a-z.A-Z0-9.-]+$"
	subject := "fake@email.com.br"

	if gorex.Test(pattern, subject) {
		fmt.Printf("/%s/ matched %s\n", pattern, subject)
	}

	posixPattern := "^[[:alnum:]_.+-]+@[[:alpha:]]+\\.[[:alnum:].-]+$"

	if gorex.Test(posixPattern, subject) {
		fmt.Printf("/%s/ matched %s\n", posixPattern, subject)
	}

}
```

## Benchmark

I trully believe that using regex.h will be more fast but I will add a benchmark anyway.

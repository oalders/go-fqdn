# go-fqdn
<!-- vim-markdown-toc GFM -->

* [About](#about)
* [Usage](#usage)

<!-- vim-markdown-toc -->

## About

This is a simple wrapper around the `net` and `os` Go standard libraries.  It
finds the Fully Qualified Domain Name of the machine the code is running on.
This is a fork of https://github.com/Showmax/go-fqdn with some API changes.

## Usage

Get the library...
```
$ go get github.com/oalders/go-fqdn
```
...and write some code.
```go
package main

import (
    "fmt"
    "log"

    fqdn "github.com/oalders/go-fqdn"
)

func main() {
    name, err := fqdn.Get()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name)
}
```

`fqdn.Get()` returns:
- machine's FQDN if found
- otherwise returns an error

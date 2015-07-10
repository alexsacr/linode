linode [![Circle CI](https://circleci.com/gh/alexsacr/linode.svg?style=shield)](https://circleci.com/gh/alexsacr/linode) [![GoDoc](https://godoc.org/github.com/alexsacr/linode?status.png)](https://godoc.org/github.com/alexsacr/linode)
======

linode is a (mostly) complete set of Go bindings to the [Linode API](https://www.linode.com/api).  It is not particularly idiomatic; the goal was to match the API as written rather than Go-ifying it.

#### Installation

Assuming you've got a working Go environment:

```sh
$ go get github.com/alexsacr/linode
```

#### Usage

About what you'd expect:

```Go
package main

import (
    "fmt"

    "github.com/alexsacr/linode"
)

func main() {
    c := linode.NewClient("yourAPIKey")

    ok := c.TestEcho()
    fmt.Println(ok)
}
```

The Linode API is fond of optional parameters, and Go hates them.

A simple solution to this conundrum would be to check if a particular parameter was zeroed or not, but that breaks down quick.  `false`, `0`, and `""` may be meaningful in a particular context.

Stealing some inspiration from the Github and DigitalOcean Go bindings, all optional parameters for API calls are pointers.  A `nil` parameter will not be sent to the API:

```Go
package main

import (
    "fmt"
    "os"

    "github.com/alexsacr/linode"
)

func main() {
    c := linode.NewClient("yourAPIKey")

    invoice, err := c.AccountEstimateInvoice("linode_new", linode.Int(1), linode.Int(2), nil)
    if err != nil {
        fmt.Printf("Oops: %s\n", err)
        os.Exit(1)
    }
    fmt.Printf("%+v\n", invoice)
}
```

Convenience functions are provided for creating pointers from literals (`linode.Int()`, `linode.String()`, and `linode.Bool()`).

API calls with a large number of optional parameters have separate option structs defined for them.

For more details, see the [godoc](http://godoc.org/github.com/alexsacr/linode).

#### Missing Methods

`avail.nodebalancers` was not added since, currently, it's a bit useless.

`linode.ip.addpublic`, with the IPv4 belt-tightening, cannot be tested and so was not added.

#### Deviations

At the time of writing, the published API has a number of errors and omissions.

`linode.disk.create` and `linode.disk.update` have multiple undocumented (but required) arguments that this package includes.

`linode.config.create` and `linode.config.update` have no documented arguments.  This package includes all (?) of the missing arguments.

The following methods have arguments that claim they accept bool values, but actually must be integers (1 or 0):
  * `pendingOnly` for `linode.job.list`
  * `isPublic` for `stackscript.create`
  * `isPublic` for `stackscript.update`
  * `checkPassive` for `nodebalancer.config.create`

`account.estimateinvoice` returns a field named `amount`, not `price`.

There is an undocumented return field called `billingmethod` in `account.info`.

`domain.update` ignores the `description` argument

`domain.list` returns an empty string if no master IPs are set, but returns `"none"` if no AXFR IPs are set.  This package returns empty strings in both cases for consistency.

#### Contributing

Pull requests are always welcome.

To get a couple tools needed for testing:

```sh
$ make setup
```

And then to run the unit tests:

```sh
$ make
```

Integration tests will cost you real money.  A few cents probably, but it's hard to say for sure.  With that warning in mind, read on.

Export an environment variable called `LINODE_API_KEY` set to, well, your API key.  The key should be for an account with nothing currently in it.  Then run:

```sh
$ make test-all
```

# go-opensea

Golang Package for Opensea API Access and Use

## Initial TODO

Here is a list of TODO's before this package is considered 'ready' for true production code.  Until then, we are in a beta stage and things will likely change.

* [ ] - error return with opensea.New
* [ ] - CICD (via github actions) for releases/tags
* [ ] - Finish structs around 'asset' response
* [ ] - v1 changes to exported CONST provided by OpenSea package
* [ ] - all package calls need to return ERR and not FATAL log in package

## Example

```go
package main

import (
    "github.com/jbkc85/go-opensea"
)

func main() {
    client = opensea.New(os.Getenv("OPENSEA_API_KEY"), "v1", "verbose")

    collection = client.GetCollection("rumble-kong-league")
    fmt.Printf("%v", collection)
}
```

# Authors

* Jason Cameron <jbkc85@gmail.com>

If you wish to support or 'say thanks' always accepting donations in ETH or Metamask Wallet currency at `0x9a8f79d5Db99e717545E49272729511Af471cA4c`

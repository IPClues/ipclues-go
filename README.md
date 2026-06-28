# ipclues-go

[![Go Reference](https://pkg.go.dev/badge/github.com/ipclues/ipclues-go.svg)](https://pkg.go.dev/github.com/ipclues/ipclues-go)
[![License: Apache License 2.0](https://img.shields.io/badge/license-Apache%20License%202.0-blue)](LICENSE)

The official Go SDK for the [IPClues](https://ipclues.com) IP intelligence API.

---

## Requirements

- Go 1.21 or later

## Installation

```bash
go get github.com/ipclues/ipclues-go
```

## Quick start

```go
package main
 
import (
    "context"
    "fmt"
    "log"
 
    "github.com/ipclues/ipclues-go/ipclues"
)
 
func main() {
    client := ipclues.New()
 
    result, err := client.LookupIP(context.Background(), "1.1.1.1")
    if err != nil {
        log.Fatal(err)
    }
 
    fmt.Println(result.IP)                    // "1.1.1.1"
    fmt.Println(result.Country.Code)          // "AU"
    fmt.Println(result.Country.Name)          // "Australia"
    fmt.Println(result.Currency.Code)         // "AUD"
    fmt.Println(result.Continent.Code)        // "OC"
}
```

# Thin wrapper for Fixer.io

This is repaired fork of [JErBerlin/gofixerio](https://github.com/JErBerlin/gofixerio), which based on [fadion/gofixerio](https://github.com/fadion/gofixerio).

## Installation

```
$ go get github.com/bedefaced/gofixerio
```

## Usage

First, let's import the package:

```go
import "github.com/bedefaced/gofixerio"
```

Let's see an exhaustive example with all the parameters:

```go
exchange := fixerio.New()
exchange.Secure(false)
exchange.AccessKey(YOUR_API_KEY)
exchange.Base(fixerio.USD) // comment line if your subscription plan at Fixer.io doesn't allow it
exchange.Symbols(fixerio.GBP, fixerio.CNY)

if response, err := exchange.GetRates(); err == nil {
    fmt.Println(response.Rates[fixerio.CNY])
}
```

Every parameter can be omitted as the package provides some sensible defaults. The base currency is `EUR`, makes a secure connection by default and returns all the supported currencies.

## Response

The response is a `ClientResponse` structure:

```go
type ClientResponse struct {
    Date  time.Time
    Base  string
    Rates rates
} 
```

`Rates` is a `map[string]float32` with currencies as keys and ratios as values. You can access them with the keys as strings or using the currency constants:

```go
response.Rates["USD"];
response.Rates[fixerio.GBP];
```

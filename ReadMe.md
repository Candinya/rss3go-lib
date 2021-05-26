# RSS3 Go Lib

> A go lib for [RSS3](https://rss3.io)

[![Go Reference](https://pkg.go.dev/badge/github.com/nyawork/rss3go_lib.svg)](https://pkg.go.dev/github.com/nyawork/rss3go_lib)

## Usage

This package includes main types of RSS3, including `RSS3Base`, `RSS3`, `RSS3List` and `RSS3Item` wrapped in `types` package.

Also, with useful methods like `*.SetSign`, and `*.CheckSign`.

Using like this:

``` go

var demo types.RSS3 // Some RSS3 Persona

demo.SetSign(key) // key: your ECDSA private key (*ecdsa.PrivateKey)

```

Can create signature for your RSS3 Persona `demo`.

``` go

demo.CheckSign()

```

Can check if the signer and embedded signature matches.

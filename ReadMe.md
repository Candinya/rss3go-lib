# RSS3 Go Lib

> A go lib for [RSS3](https://rss3.io)

[![Go Reference](https://pkg.go.dev/badge/github.com/nyawork/rss3go-lib.svg)](https://pkg.go.dev/github.com/nyawork/rss3go-lib)

## Usage

### Types

This package includes main types of RSS3, including:

- RSS3Base
- RSS3
- RSS3List
- RSS3Item

Also, with useful methods like `*.SetSign`, and `*.CheckSign`.

Using like this can create signature for your RSS3 Persona `demo`.

``` go

var key *ecdsa.PrivateKey

...

var demo types.RSS3

demo.SetSign(key)

```

And this can check if the signer and embedded signature matches.

``` go

demo.CheckSign()

```

For other demos, you can check those `*_test.go` files for more help.

### Utils

Contains some useful tools:

- FixSign
- SignMsg

#### FixSign

This function can delete the starting "0x" of signature (if it has), and fix the ending (eth-crypto library would somewhat add 0x1b to the last byte, causing it change from 00/01 to 1b/1c, which will crash the requested functions)

``` go

utils.FixSign(signatureReceived)

```

#### SignMsg

This function can be used to create signatures for `[]byte` . Here's a demo with a string.

``` go

var key *ecdsa.PrivateKey

...

msgBytes := []byte("DemoMsg")

sign := utils.SignMsg(key, msgBytes)

signHex := hex.EncodeToString(sign)

```

## Inner Types 

> (Recommend not to use outside this package)

- types/for_sign.go: types derived of useless fields for creating signature.
  
  Those types often ends with `4S` which means "For(**4**) **S**ign"
  
- types/for_file.go: full types with all fields (as described in definition) for storage use.

  Those types often ends with `4F` which means "For(**4**) **F**ile"

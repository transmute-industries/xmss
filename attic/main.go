package main

import (
    "fmt"
    "crypto/rand"
    "example.com/sanity/keypair"
)


func main() {

    // generate
    jwk, err1 := keypair.Generate()
    fmt.Println(jwk)
    if err1 != nil {
        fmt.Println(err1)
    }

    // sign
    msg := make([]byte, 32)
	rand.Read(msg)
    sig, err2 := keypair.Sign(jwk, msg)
    fmt.Println(sig)
    if err2 != nil {
        fmt.Println(err2)
    }

    // verify
    v, err3 := keypair.Verify(msg, sig, jwk)
    fmt.Println(v)
    if err3 != nil {
        fmt.Println(err3)
    }
}

package main

import (
    "fmt"
    "github.com/sangyun-han/strongswan-go/vici"
)

func main() {
    fmt.Println("test")
    client, err := vici.NewClientConnFromDefaultSocket()
    if err != nil {
        panic(err)
    }
    defer client.Close()

    v, err := client.Version()
    if err != nil {
        panic(err)
    }

    fmt.Printf("%#v \n", v)

}

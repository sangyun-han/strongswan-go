package main

import (

    "github.com/bronze1man/goStrongswanVici"
    "fmt"
)

func main() {
    fmt.Println("test")
    client, err := goStrongswanVici.NewClientConnFromDefaultSocket()
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

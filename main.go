// main.go
package main

import (
    "fmt"
    "os"
    "training/data"
    "training/function"
)

func main() {
    args := os.Args[1:]

    if len(args) < 1 {
        fmt.Println("Argumen Kosong")
        return
    }

    namaTeman := args[0]

    function.TampilkanDataTeman(namaTeman, data.TemanList)
}

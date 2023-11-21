// function/tampilan.go
package function

import (
    "fmt"
    "training/method"
)

func TampilkanDataTeman(nama string, temanList []method.Teman) {
    for _, teman := range temanList {
        if teman.Nama == nama {
            teman.TampilkanData()
            return
        }
    }
    fmt.Printf("Tidak ada data teman dengan nama %s\n", nama)
}

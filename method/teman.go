// method/teman.go
package method

import "fmt"

type Teman struct {
    ID        int
    Nama      string
    Alamat    string
    Pekerjaan string
    Alasan    string
}

func (t Teman) TampilkanData() {
    fmt.Printf("Data Teman:\nAbsen: %d\nNama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n", t.ID, t.Nama, t.Alamat, t.Pekerjaan, t.Alasan)
}

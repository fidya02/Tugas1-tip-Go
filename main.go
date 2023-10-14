package main

import 
	"fmt"

func main() { 
	var tagihan int
	fmt.Print("Masukkan jumlah tagihan: ")
	fmt.Scanln(&tagihan)

	// Untuk menghitung tip
    var tip float64

	// Untuk menghitung tip berdasarkan kondisi
    if tagihan >= 50 && tagihan <= 300 {
        tip = float64(tagihan) * 0.15
    } else {
        tip = float64(tagihan) * 0.20
    }

    // Menghitung total (tagihan) + tip)
    total := float64(tagihan) + tip

    // Menampilkan hasil pada konsol
    fmt.Printf("tagihannya %d, tipnya %.2f, dan total nilainya %.2f\n", tagihan, tip, total)
}

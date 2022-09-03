package main

import "fmt"

type produk struct {
	nama  string
	harga int
}

//mencari harga 10.000
func mencariHarga(prod []produk, price int) []string {
	//Find matching value(s)
	var hasil []string
	for i := 0; i < len(prod); i++ {
		if prod[i].harga == price {
			hasil = append(hasil, prod[i].nama)
		}
	}
	return hasil
}

//mencari produk murah dan mahal
func mencariMinMax(prod []produk) (produk, produk) {
	for i := 0; i < len(prod); i++ {
		for j := 1; j < len(prod); j++ {
			if prod[i].harga > prod[j].harga {
				temp := prod[i]
				prod[i] = prod[j]
				prod[j] = temp
			}
		}
	}
	return prod[0], prod[len(prod)-1]
}

func main() {
	var listProduk = [10]produk{
		{"Benih Lele", 50000},
		{"Pakan lele cap menara", 25000},
		{"Probiotik A", 75000},
		{"Probiotik Nila B", 10000},
		{"Pakan Nila", 20000},
		{"Benih Nila 1", 20000},
		{"Cupang", 5000},
		{"Benih Nila 2", 30000},
		{"Benih Cupang", 10000},
		{"Probiotik B", 10000},
	}

	//produk dengan harga Rp 10.000
	daftar := mencariHarga(listProduk[:], 10000)
	fmt.Println("Daftar produk dengan harga Rp 10.000 :")
	for _, value := range daftar {
		fmt.Println("- " + value)
	}

	fmt.Println("------------------------------------------------")

	//produk termurah dan termahal
	prodMur, prodMah := mencariMinMax(listProduk[:])
	fmt.Println("Daftar produk termurah: \n", prodMur.nama, " Rp ", prodMur.harga)
	fmt.Println("Daftar produk termahal: \n", prodMah.nama, " Rp ", prodMah.harga)

	fmt.Println("------------------------------------------------")

	//pembelian total Rp 100.000
	fmt.Println("Berikut total pembelian produk hingga Rp 100.000: ")
	//pembelian Rp 100.000
	jumlah_total := 0
	for i := 0; i < len(listProduk); i++ {
		if jumlah_total >= 50000 && jumlah_total <= 110000 {
			break
		} else {
			fmt.Println(listProduk[i].nama, " Rp ", listProduk[i].harga)
			jumlah_total = listProduk[i].harga + jumlah_total
		}
	}
	fmt.Println("Total     : ", jumlah_total)
}

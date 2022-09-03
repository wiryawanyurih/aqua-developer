package main

import "fmt"

//struct section
type lecture struct {
	name    string
	address string
	matkul  string
	boot    academy
}

type academy struct {
	name    string
	address string
}

func main() {
	//function struct section
	var s1 lecture
	s1.name = "Wiryawan"
	s1.address = "Jl. Suhat"
	s1.matkul = "Golang"
	s1.boot = academy{
		name:    "eFishery Academy",
		address: "Jl. Kayu Tangan",
	}

	//edit struct section
	s1.address = "Jl. Pasar Besar"

	//print struct section
	fmt.Println("Nama Pengajar   : ", s1.name)
	fmt.Println("Alamat Pengajar : ", s1.address)
	fmt.Println("Mata Kuliah     : ", s1.matkul)
	fmt.Println("Akademi         : ", s1.boot.name)
	fmt.Println("Alamat Akademi  : ", s1.boot.address)

	//data structure - map
	var kampus map[string]academy
	kampus = map[string]academy{}

	kampus["kampus_pusat"] = academy{
		name:    "Pusat",
		address: "Jl. Kayu tanggan, Kota Malang",
	}

	//edit map section
	kampus["kampus_pusat"] = academy{
		name:    "pusat",
		address: "Jl. Bandung",
	}

	//print map section
	fmt.Println("Posisi Kampus : ", kampus["kampus_pusat"].name)
	fmt.Println("Alamat Kampus : ", kampus["kampus_pusat"].address)

	//array section
	var daftar_matkul = [3]string{"Golang", "PHP", "Java"}

	//edit array section
	daftar_matkul[1] = "Mysql"

	//print array section
	for _, v := range daftar_matkul {
		fmt.Println("Mata Kuliah: ", v)
	}

	//slice section
	var kurikulum = []string{"2018", "2019", "2020"}

	//edit slice
	kurikulum[1] = "2021"

	//print slice section
	fmt.Println(kurikulum[0:2])
}

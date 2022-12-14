package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	Nama, Alamat, Pekerjaan, Alasan string
}

func main() {

	var data = []Biodata{
		{
			Nama:      "Fitri",
			Alamat:    "Jl. Lorem",
			Pekerjaan: "Backend",
			Alasan:    "Alasan Fitri",
		},
		{
			Nama:      "Ayu",
			Alamat:    "Jl. Ipsum",
			Pekerjaan: "Frontend",
			Alasan:    "Alasan Ayu",
		},
		{
			Nama:      "Anggraini",
			Alamat:    "Jl. Dolor",
			Pekerjaan: "Fullstack",
			Alasan:    "Alasan Anggraini",
		},
		{
			Nama:      "Fiyuang",
			Alamat:    "Jl. Sit",
			Pekerjaan: "DevOps",
			Alasan:    "Alasan Fiyuang",
		},
		{
			Nama:      "Fifiyuu",
			Alamat:    "Jl. Amet",
			Pekerjaan: "Mobile Dev",
			Alasan:    "Alasan Fifiyuu",
		},
	}

	peserta := os.Args
	if len(peserta) < 2 {
		fmt.Println("Tolong masukan nama atau nomor absen")
		fmt.Println("Contoh: 'go run main.go Fitri' atau 'go run main.go 2'")
	} else {
		// check args a number or string
		if key, err := strconv.Atoi(peserta[1]); err == nil {
			if key < len(data) {
				printDetailBiodata(data[key], key)
			} else {
				fmt.Println("Tidak ada absen dengan nomor yang anda input, silahkan input kembali")
			}
		} else {
			for key, value := range data {
				if value.Nama == peserta[1] {
					printDetailBiodata(data[key], key)
				}
			}
		}
	}

}

func printDetailBiodata(class Biodata, key int) {
	fmt.Println("ID :", key)
	fmt.Println("nama :", class.Nama)
	fmt.Println("alamat :", class.Alamat)
	fmt.Println("pekerjaan :", class.Pekerjaan)
	fmt.Println("alasan :", class.Alasan)
}

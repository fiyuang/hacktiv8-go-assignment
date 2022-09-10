package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
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
	}

	peserta := os.Args
	for key, value := range data {
		if value.Nama == peserta[1] {
			fmt.Println("ID :", key)
			fmt.Println("nama :", data[key].Nama)
			fmt.Println("alamat :", data[key].Alamat)
			fmt.Println("pekerjaan :", data[key].Pekerjaan)
			fmt.Println("alasan :", data[key].Alasan)
		}
	}
}

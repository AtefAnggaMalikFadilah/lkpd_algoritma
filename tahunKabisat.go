package main

import "fmt"

func main(){

	tahunAkhir := 2024

	for tahunAwal := 1600; tahunAwal <= tahunAkhir; tahunAwal++ {
		if tahunAwal % 400 == 0{
			fmt.Print(tahunAwal, " adalah tahun kabisat\n")
		} else if tahunAwal % 100 == 0{
			fmt.Print(tahunAwal, " bukan tahun kabisat\n")
		} else if tahunAwal % 4 == 0 {
			fmt.Print(tahunAwal, " adalah tahun kabisat\n")
		} else {
			fmt.Print("")
		}
	}
}
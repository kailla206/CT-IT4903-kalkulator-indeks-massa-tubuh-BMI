package main

import (
	"fmt"
)

func main() {
	for {
		var nama string
		var berat, tinggi float64

		fmt.Println("\n=== Program Perhitungan BMI ===")

		// Input nama
		fmt.Print("Masukkan nama Anda: ")
		fmt.Scanln(&nama)

		// Validasi berat badan
		for {
			fmt.Print("Masukkan berat badan (kg): ")
			fmt.Scanln(&berat)
			if berat > 0 {
				break
			}
			fmt.Println("ERROR: Berat badan harus lebih dari 0.")
		}

		// Validasi tinggi badan
		for {
			fmt.Print("Masukkan tinggi badan (m): ")
			fmt.Scanln(&tinggi)
			if tinggi > 0 {
				break
			}
			fmt.Println("ERROR: Tinggi badan harus lebih dari 0.")
		}

		// Hitung BMI
		bmi := berat / (tinggi * tinggi)

		// Tentukan kategori
		var kategori string
		if bmi < 18.5 {
			kategori = "Kurus"
		} else if bmi < 25 {
			kategori = "Normal"
		} else if bmi < 30 {
			kategori = "Gemuk"
		} else {
			kategori = "Obesitas"
		}

		// Hitung berat badan ideal (BMI 22)
		beratIdeal := 22 * (tinggi * tinggi)

		// Output
		fmt.Println("\n=== Hasil Perhitungan ===")
		fmt.Println("Nama            :", nama)
		fmt.Printf("BMI             : %.2f\n", bmi)
		fmt.Println("Kategori        :", kategori)
		fmt.Printf("Berat Ideal     : %.2f kg\n", beratIdeal)

		// Menu ulang
		var ulang string
		fmt.Print("\nHitung ulang? (y/n): ")
		fmt.Scanln(&ulang)

		if ulang != "y" && ulang != "Y" {
			fmt.Println("Program selesai. Terima kasih!")
			break
		}
	}
}


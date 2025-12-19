package main

import (
	"fmt"
)

func main() {
	ulang := true
	for ulang {

		var nama, opsi string
		var berat, tinggi float64
		valid := false
		
		fmt.Println("\n=== Program Perhitungan BMI ===")

		fmt.Print("Masukkan nama pengguna: ")
		fmt.Scanln(&nama)

		fmt.Print("Masukkan berat badan (kg): ")
		fmt.Scanln(&berat)

		// validasi tinggi badan
		for valid == false {
			fmt.Print("Masukkan tinggi badan (m): ")
			fmt.Scanln(&tinggi)

			if tinggi > 0 {
				valid = true
			} else {
				fmt.Println("Tinggi badan harus lebih dari 0.")
			}
		}

		// Menghitung indeks massa tubuh (BMI)
		bmi := berat / (tinggi * tinggi)

		// Berat badan ideal (BMI ideal=22)
		beratIdeal := 22 * (tinggi * tinggi)

		// Menentukan kategori BMI
		var kategori string
		if bmi < 18.5 {
			kategori = "Kurus"
		} else if bmi < 25 {
			kategori = "Normal"
		} else {
			kategori = "Gemuk"
		}

		// Output
		fmt.Println("\n=== Hasil ===")
		fmt.Println("Nama:", nama)
		fmt.Printf("BMI: %.2f\n", bmi)
		fmt.Println("Kategori:", kategori)
		fmt.Printf("Berat Ideal: %.2f kg\n", beratIdeal)

		// Saran penambahan atau pengurangan berat badan
		if kategori == "Kurus" {
			fmt.Printf("Saran: Tambah berat ± %.2f kg\n", beratIdeal - berat)
		}

		if kategori == "Gemuk" {
			fmt.Printf("Saran: Kurangi berat ± %.2f kg\n", berat - beratIdeal)
			}

			tips := []string{
				"Kurangi makanan berlemak dan kadar gula yang tinggi",
				"perbanyak konsumsi sayur dan buah",
				"Rutin berolahraga minimal 30 menit per hari",
				"Pola tidur yang teratur",
				"Minum air mineralyang cukup"
		}
			fmt.Println("Tips penurunan berat badan: ")
			for i := 0; i < len(tips); i++ {
				fmt.Printf("-%s\n", tips[i])
			}
		}

		// opsi pengulangan
		fmt.Print("\nHitung lagi? (ya/tidak): ")
		fmt.Scanln(&opsi)

		if opsi != "ya" {
			fmt.Println("Selesai, Terima kasih! ")
			ulang = false
		}
	}
}









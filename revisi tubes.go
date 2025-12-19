package main

import (
	"fmt"
)

func main() {
	ulang := true
	for ulang {

		var nama string
		var berat, tinggi float64
		valid := false
		var pilihan string

		fmt.Println("\n=== Program Perhitungan BMI ===")

		fmt.Print("Masukkan nama pengguna: ")
		fmt.Scanln(&nama)

		fmt.Print("Masukkan berat badan (kg): ")
		fmt.Scanln(&berat)

		// while-loop validasi tinggi badan
		for valid == false {
			fmt.Print("Masukkan tinggi badan (m): ")
			fmt.Scanln(&tinggi)

			if tinggi > 0 {
				valid = true
			} else {
				fmt.Println("Tinggi badan harus lebih dari 0.")
			}
		}

		// Hitung BMI
		bmi := berat / (tinggi * tinggi)

		// Berat badan ideal (BMI 22)
		beratIdeal := 22 * (tinggi * tinggi)

		// Kategori BMI
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

		// Saran berat badan
		if kategori == "Kurus" {
			fmt.Printf("Saran: Tambah berat ± %.2f kg\n", beratIdeal-berat)
		}

		if kategori == "Gemuk" {
			fmt.Printf("Saran: Kurangi berat ± %.2f kg\n", berat-beratIdeal)
		}

		// while-loop kontrol pengulangan
		fmt.Print("\nHitung lagi? (ya/tidak): ")
		fmt.Scanln(&pilihan)

		if pilihan != "ya" {
			fmt.Println("Selesai, Terima kasih! ")
			ulang = false
		}
	}
}



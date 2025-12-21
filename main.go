package main

import (
	"fmt"
)

func main() {
	ulang := true
	for ulang {

		var nama string
		var berat, tinggi float64
		
		fmt.Println("\n=== Program Perhitungan BMI ===")

		fmt.Print("Masukkan nama pengguna: ")
		fmt.Scanln(&nama)

		// Validasi berat dan tinggi badan
		validBerat := false
		for validBerat == false {
			fmt.Print("Masukkan berat badan (kg): ")
			n, _ := fmt.Scanln(&berat)

			if n != 1 {
				fmt.Println("Input harus berupa angka ")
			} else if berat <= 0 {
				fmt.Println("Berat badan harus lebih dari 0.")
			} else {
				validBerat = true
			}
		}
		validTinggi := false
		for validTinggi == false {
			fmt.Print("Masukkan tinggi badan (m): ")
			n, _ := fmt.Scanln(&tinggi)

			if n != 1 {
				fmt.Println("Input harus berupa angka ")
			} else if tinggi <= 0 {
				fmt.Println("Tinggi badan harus lebih dari 0.")
			} else {
				validTinggi = true
			}
		}

		// Menghitung indeks massa tubuh (BMI)
		bmi := berat / (tinggi * tinggi)

		// Berat badan ideal (BMI ideal=22)
		beratIdeal := 22 * (tinggi * tinggi)

		// Menentukan kategori BMI
		var kategori string
		if 0 < bmi && bmi < 18.5 {
			kategori = "Kurus"
		} else if 18.6 < bmi && bmi < 25 {
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

			tips:= []string{
				"Pilih makanan sehat yang kaya akan nutrisi",
				"Makan lebih sering dengan porsi kecil", 
				"Konsumsi minuman padat kalori dan nutrisi",
				"Tambah asupan kalori",
				"Rutin berolahraga",
		}
			fmt.Println("Tips penambahan berat badan: ")
			for i := 0; i < len(tips); i++ {
				fmt.Printf("-%s\n", tips[i])
			}
		}

		if kategori == "Gemuk" {
			fmt.Printf("Saran: Kurangi berat ± %.2f kg\n", berat - beratIdeal)

			tips := []string{
				"Kurangi makanan berlemak dan kadar gula yang tinggi",
				"Perbanyak konsumsi sayur dan buah",
				"Rutin berolahraga minimal 30 menit per hari",
				"Pola tidur yang teratur",
				"Minum air mineral yang cukup",
			}
		
			fmt.Println("Tips penurunan berat badan: ")
			for i := 0; i < len(tips); i++ {
				fmt.Printf("-%s\n", tips[i])
			}
		}

		// Opsi pengulangan
		var opsi string
		fmt.Print("\nApakah tips ini membantu? (ya/tidak): ")
		fmt.Scanln(&opsi)

		if opsi != "tidak" {
			fmt.Println("Terima kasih, senang bisa membantu! ")
			ulang = false
		} else {
			fmt.Println("Mohon maaf, kami akan berusaha untuk meningkatkan layanan kami kedepannya ")
			ulang = false
		}
	}
}
















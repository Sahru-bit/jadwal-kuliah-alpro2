package main

import (
	"fmt"
	"strings"
)

const NMAX int = 10

type jadwal struct {
	hari    string
	mulai   string
	selesai string
	matkul  string
}

type tabJadwal [NMAX]jadwal

func cekKuliah(T tabJadwal, n int, hari, jam string) (jadwal, bool) {
	var i int

	for i = 0; i < n; i++ {
		if T[i].hari == hari &&
			jam >= T[i].mulai &&
			jam <= T[i].selesai {

			return T[i], true
		}
	}

	return jadwal{}, false
}

func tampilJadwalHari(T tabJadwal, n int, hari string) {
	var i int
	var ada bool

	ada = false

	fmt.Println("\nJadwal Hari", hari)
	fmt.Println("--------------------------")

	for i = 0; i < n; i++ {
		if T[i].hari == hari {
			fmt.Println(T[i].mulai, "-", T[i].selesai, ":", T[i].matkul)
			ada = true
		}
	}

	if !ada {
		fmt.Println("Tidak ada jadwal pada hari tersebut")
	}
}

func main() {
	var data tabJadwal
	var hari, jam string
	var hasil jadwal
	var ketemu bool
	var pilihan int
	var n int

	n = 10

	data[0].hari = "SENIN"
	data[0].mulai = "08:30"
	data[0].selesai = "10:30"
	data[0].matkul = "ALGORITMA DAN PEMROGRAMAN 2"

	data[1].hari = "SENIN"
	data[1].mulai = "11:30"
	data[1].selesai = "13:30"
	data[1].matkul = "ETIKA DALAM AI"

	data[2].hari = "SELASA"
	data[2].mulai = "07:30"
	data[2].selesai = "10:30"
	data[2].matkul = "PEMODELAN BASIS DATA"

	data[3].hari = "RABU"
	data[3].mulai = "08:30"
	data[3].selesai = "10:30"
	data[3].matkul = "BAHASA INGGRIS"

	data[4].hari = "RABU"
	data[4].mulai = "11:30"
	data[4].selesai = "14:30"
	data[4].matkul = "ORGANISASI DAN ARSITEKTUR KOMPUTER"

	data[5].hari = "KAMIS"
	data[5].mulai = "10:30"
	data[5].selesai = "12:30"
	data[5].matkul = "ALGORITMA DAN PEMROGRAMAN 2"

	data[6].hari = "JUMAT"
	data[6].mulai = "08:30"
	data[6].selesai = "10:30"
	data[6].matkul = "KALKULUS LANJUT"

	data[7].hari = "JUMAT"
	data[7].mulai = "13:30"
	data[7].selesai = "16:30"
	data[7].matkul = "MATRIKS DAN RUANG VEKTOR"

	data[8].hari = "SABTU"
	data[8].mulai = "10:30"
	data[8].selesai = "12:30"
	data[8].matkul = "KALKULUS LANJUT"

	data[9].hari = "SABTU"
	data[9].mulai = "13:30"
	data[9].selesai = "16:30"
	data[9].matkul = "ALGORITMA DAN PEMROGRAMAN 2 (Praktikum)"

	fmt.Println("===== MENU =====")
	fmt.Println("1. Cek kuliah yang sedang berlangsung")
	fmt.Println("2. Tampilkan semua jadwal pada hari tertentu")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {

		fmt.Print("Masukkan hari : ")
		fmt.Scan(&hari)

		hari = strings.ToUpper(hari)

		fmt.Print("Masukkan jam (HH:MM) : ")
		fmt.Scan(&jam)

		hasil, ketemu = cekKuliah(data, n, hari, jam)

		if !ketemu {
			fmt.Println("Tidak ada jadwal perkuliahan")
		} else {
			fmt.Println("Perkuliahan sedang berlangsung :", hasil.matkul)
			fmt.Println("Jam perkuliahan               :", hasil.mulai, "-", hasil.selesai)
		}

	} else if pilihan == 2 {

		fmt.Print("Masukkan hari : ")
		fmt.Scan(&hari)

		hari = strings.ToUpper(hari)

		tampilJadwalHari(data, n, hari)

	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

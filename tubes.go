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

	data[0] = jadwal{"SENIN", "08:30", "10:30", "ALGORITMA DAN PEMROGRAMAN 2"}
	data[1] = jadwal{"SENIN", "11:30", "13:30", "ETIKA DALAM AI"}
	data[2] = jadwal{"SELASA", "07:30", "10:30", "PEMODELAN BASIS DATA"}
	data[3] = jadwal{"RABU", "08:30", "10:30", "BAHASA INGGRIS"}
	data[4] = jadwal{"RABU", "11:30", "14:30", "ORGANISASI DAN ARSITEKTUR KOMPUTER"}
	data[5] = jadwal{"KAMIS", "10:30", "12:30", "ALGORITMA DAN PEMROGRAMAN 2"}
	data[6] = jadwal{"JUMAT", "08:30", "10:30", "KALKULUS LANJUT"}
	data[7] = jadwal{"JUMAT", "13:30", "16:30", "MATRIKS DAN RUANG VEKTOR"}
	data[8] = jadwal{"SABTU", "10:30", "12:30", "KALKULUS LANJUT"}
	data[9] = jadwal{"SABTU", "13:30", "16:30", "ALGORITMA DAN PEMROGRAMAN 2 (Praktikum)"}

	fmt.Println("===== MENU =====")
	fmt.Println("1. Cek kuliah yang sedang berlangsung")
	fmt.Println("2. Tampilkan semua jadwal pada hari tertentu")
	fmt.Println("3. Tampilkan semua jadwal kuliah")
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

	} else if pilihan == 3 {

		tampilJadwalHari(data, n, "SENIN")
		tampilJadwalHari(data, n, "SELASA")
		tampilJadwalHari(data, n, "RABU")
		tampilJadwalHari(data, n, "KAMIS")
		tampilJadwalHari(data, n, "JUMAT")
		tampilJadwalHari(data, n, "SABTU")

	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

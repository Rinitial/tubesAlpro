package main

import (
	"fmt"
)

const NMAX = 10

type dataPart struct {
	spare   string
	harga   float64
	terjual int
}

type arrParts [NMAX]dataPart
type parts [NMAX]string

type dataPengguna struct {
	nama         string
	bulanService int
	parts        parts
	total        float64
}
type arrPengguna [NMAX]dataPengguna

func main() {
	var m, n int
	var dataPengguna arrPengguna
	var dataParts arrParts
	inputParts(&dataParts, &m)
	menu(&dataPengguna, &dataParts, &n, m)
}

// I.S. array y merupakan array yang terdiri dari spare dan harga yang akan di input
// F.S. array y terisi data sembarang
func inputParts(y *arrParts, m *int) {
	fmt.Print("Masukkan Berapa Parts yang ingin diinput : ")
	fmt.Scan(m)
	for i := 0; i < *m; i++ {
		fmt.Println("Masukkan Nama Parts : ")
		fmt.Scan(&y[i].spare)
		fmt.Println("Masukkan Harga Parts : ")
		fmt.Scan(&y[i].harga)
	}
}

// I.S. menu menunggu pilihan masukkan dari user
// F.S. menampilan pilihan menu dan mengarahkan user
func menu(x *arrPengguna, y *arrParts, n *int, m int) {
	var pilihan int
	var berdasarkan int
	for pilihan != 6 {
		fmt.Println("╒════════════════════════════════╕")
		fmt.Println("│           Pilih Menu           │")
		fmt.Println("├────────────────────────────────┤")
		fmt.Println("│ 1. Tambah Data                 │")
		fmt.Println("│ 2. Edit Data                   │")
		fmt.Println("│ 3. Hapus Data                  │")
		fmt.Println("│ 4. Tampilkan Data              │")
		fmt.Println("│ 5. Tampilkan Parts             │")
		fmt.Println("│ 6. Cari Data Pengguna          │")
		fmt.Println("│ 7. Keluar                      │")
		fmt.Println("╘════════════════════════════════╛")
		fmt.Print("Masukkan menu: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahData(x, y, n, m)
		case 2:
			editData(x, y, *n, m)
		case 3:
			hapusData(x, y, n, m)
		case 4:
			tampilData(*x, n)
			fmt.Println("╒════════════════════════════════╕")
			fmt.Println("│           Pilih Menu           │")
			fmt.Println("├────────────────────────────────┤")
			fmt.Println("│ 1. Urut berdasarkan bulan      │")
			fmt.Println("│    terbaru                     │")
			fmt.Println("│ 2. Urut berdasarkan bulan      │")
			fmt.Println("│    terlama                     │")
			fmt.Println("│ 3. urutkan berdasarkan spare   │")
			fmt.Println("│    terlaris                    │")
			fmt.Println("│ 4. kembali ke menu             │")
			fmt.Println("╘════════════════════════════════╛")
			fmt.Print("Masukkan menu: ")
			fmt.Scan(&berdasarkan)
			switch berdasarkan {
			case 1:
				menaikBulan(x, n)
			case 2:
				menurunBulan(x, n)
			case 3:
				SpareTerlaris(y, m)
			case 4:
			}
		case 5:
			tampilParts(*y, m)
		case 6:
			binSearch(x, n)
		case 7:
			fmt.Println("Keluar dari program.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// I.S. array x merupakan array yang terdiri dari nama, bulan, parts, dan total
// F.S. arrray x terisi dengan data sembarang dan dilakukkan update parts terjual pada array y dan dilakukkan total harga pada array yang diambil data harga pada array y
func tambahData(x *arrPengguna, y *arrParts, n *int, m int) {
	var spare parts

	fmt.Println("----------------")
	fmt.Printf("Data Pelanggan Ke - %d\n", *n+1)
	fmt.Println("Masukkan Nama:")
	fmt.Scan(&x[*n].nama)
	fmt.Println("Masukkan Bulan Service: (1 - 12)")
	fmt.Scan(&x[*n].bulanService)
	for x[*n].bulanService < 1 || x[*n].bulanService > 12 {
		fmt.Println("Masukkan Bulan Service: (1 - 12)")
		fmt.Scan(&x[*n].bulanService)
	}
	for i := 1; i <= NMAX && spare[i-1] != "none"; i++ {
		fmt.Println("Masukkan Spare Part : (Masukkan 'none' jika sudah)")
		fmt.Scan(&spare[i])
		if spare[i] != "none" {
			for j := 0; j < m; j++ {
				if spare[i] == y[j].spare {
					y[j].terjual = y[j].terjual + 1
					x[*n].total += y[j].harga
				}
			}
			x[*n].parts = spare
		}
	}
	*n++
}

// I.S. terdefinisi array x dan y berisi data pelanggan dan spare parts
// F.S. data array x teredit berdasarkan pencarian dengan nama dan bulan dan melakukan update data terjual pada array y
func editData(x *arrPengguna, y *arrParts, n int, m int) {
	var spare parts
	if n == 0 {
		fmt.Println("Data Kosong")
		fmt.Println("Silahkan masukkan data terlebih dahulu")
	}

	fmt.Println("EDIT DATA:")

	index := findIndexSeq(*x, n)
	x[index].total = 0
	// mengurangi jumlah terjual dari database parts
	for i := 0; i < len(x[index].parts); i++ {
		for j := 0; j < m; j++ {
			if x[index].parts[i] == y[j].spare {
				y[j].terjual = y[j].terjual - 1
			}
		}
	}

	if index != -1 {
		for i := 1; i <= len(x[index].parts) && spare[i-1] != "none"; i++ { // Modify the loop condition
			fmt.Println("Masukkan Spare Part : (Masukkan 'none' jika sudah) ")
			fmt.Scan(&spare[i])
			if spare[i] != "none" {
				for j := 0; j < m; j++ {
					if spare[i] == y[j].spare {
						y[j].terjual = y[j].terjual + 1
						x[index].total += y[j].harga
					}
				}
				x[index].parts = spare
			}
		}
	}
}

// I.S. terdefinisi array x dan y berisi data pelanggan dan spare parts
// F.S. data array x terhapus berdasarkan pencarian dengan nama dan bulan dan melakukan update data terjual pada array y
func hapusData(x *arrPengguna, y *arrParts, n *int, m int) {

	if *n == 0 {
		fmt.Println("Data Kosong")
		fmt.Println("Silahkan masukkan data terlebih dahulu")
	}

	var nama string
	var bulan int

	fmt.Println("HAPUS DATA :")

	fmt.Println("Masukkan nama dan bulan service pelanggan yang ingin dihapus :")
	fmt.Scan(&nama, &bulan)

	for i := 0; i < *n; i++ {

		if x[i].nama == nama && x[i].bulanService == bulan {
			// mengurangi jumlah terjual dari database parts
			for j := 0; j < len(x[i].parts); j++ {
				for k := 0; k < m; k++ {
					if x[i].parts[j] == y[k].spare {
						y[k].terjual = y[k].terjual - 1
					}
				}
			}

			for j := i; j < *n-1; j++ {
				x[j] = x[j+1]

			}
			*n = *n - 1
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	}

}

// mengembalikkan index yang dicari berdasarkan nama bulan yang cocok dengan sequential search
func findIndexSeq(x arrPengguna, n int) int {
	var nama string
	var bulan int
	index := -1
	fmt.Println("Masukkan data nama dan bulan pelanggan servis yang ingin diedit:")
	fmt.Scan(&nama, &bulan)
	for i := 0; i < n && index == -1; i++ {
		if x[i].nama == nama && x[i].bulanService == bulan {
			index = i
		}
	}
	return index
}

// mengembalikkan index yang dicari berdasarkan nama bulan yang cocok dengan binary search
func findIndexBin(x arrPengguna, n int) int {
	var nama string
	var bulan, ki, ka, te int
	ki = 0
	ka = n - 1
	index := -1
	fmt.Println("Masukkan data nama dan bulan pelanggan servis yang ingin dilihat:")
	fmt.Scan(&nama, &bulan)
	for ki <= ka && index == -1 {
		te = (ki + ka) / 2
		if bulan > x[te].bulanService && nama == x[te].nama {
			ka = te - 1
		} else if bulan > x[te].bulanService && nama == x[te].nama {
			ki = te + 1
		} else {
			index = te
		}
	}
	return index
}

// I.S. terdefinisi data array x yang berisi sebuah data pelanggan
// F.S. melakuan sorting lalu dilakukan binary search dan menampilkan isi data yang dicari berdasarkan nama dan bulannya
func binSearch(x *arrPengguna, n *int) {
	menurunBulan(x, n)
	index := findIndexBin(*x, *n)
	if index == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println("Data ditemukan")
		fmt.Println("╒════════════════════════════════╕")
		fmt.Println("│           Data History         │")
		fmt.Println("├────────────────────────────────┤")
		fmt.Printf("│%v. Nama : %-22s│\n", index+1, x[index].nama)
		fmt.Printf("│   Bulan: %-22v│\n", x[index].bulanService)
		for j := 1; j <= len(x[index].parts[j]); j++ {
			fmt.Printf("│   Parts: %-22s│\n", x[index].parts[j])
		}
		fmt.Printf("│   Total: %-22v│", x[index].total)
		fmt.Println()
	}
	fmt.Println("╘════════════════════════════════╛")
}

// I.S. terdefinisi data array x yang berisi sebuah data pelanggan
// F.S. menampilkan hasil isi dari array x
func tampilData(x arrPengguna, n *int) {
	k := 1
	fmt.Println("╒════════════════════════════════╕")
	fmt.Println("│           Data History         │")
	fmt.Println("├────────────────────────────────┤")
	for i := 0; i < *n; i++ {
		fmt.Printf("│%v. Nama : %-22s│\n", k, x[i].nama)
		fmt.Printf("│   Bulan: %-22v│\n", x[i].bulanService)
		for j := 1; j <= len(x[i].parts[j]); j++ {
			fmt.Printf("│   Parts: %-22s│\n", x[i].parts[j])
		}
		fmt.Printf("│   Total: %-22v│", x[i].total)
		fmt.Println()
		k++
	}
	fmt.Println("╘════════════════════════════════╛")
}

// I.S. terdefinisi data array y yang berisi sebuah data spare parts
// F.S. menampilkan hasil isi dari array y
func tampilParts(y arrParts, m int) {
	k := 1
	fmt.Println("Data Parts : ")
	fmt.Println("╒════════════════════════════════╕")
	fmt.Println("│           Data Parts           │")
	fmt.Println("├────────────────────────────────┤")
	for i := 0; i < m; i++ {
		fmt.Printf("│%v. parts: %-22s│\n", k, y[i].spare)
		fmt.Printf("│   Harga: %-22v│\n", y[i].harga)
		fmt.Printf("│   Total: %-22v│", y[i].terjual)
		fmt.Println()
		k++
	}
	fmt.Println("╘════════════════════════════════╛")
}

// I.S. terdefinisi data array x yang berisi sebuah data pelanggan
// F.S. menampilkan hasil insertion sort
func menaikBulan(x *arrPengguna, n *int) {
	for i := 1; i < *n; i++ {
		key := x[i]
		j := i - 1
		for j >= 0 && x[j].bulanService > key.bulanService {
			x[j+1] = x[j]
			j--
		}
		x[j+1] = key
	}
	k := 1
	fmt.Println("╒════════════════════════════════╕")
	fmt.Println("│   Data History Bulan Terbaru   │")
	fmt.Println("├────────────────────────────────┤")
	for i := 0; i < *n; i++ {
		fmt.Printf("│%v. Nama : %-22s│\n", k, x[i].nama)
		fmt.Printf("│   Bulan: %-22v│\n", x[i].bulanService)
		for j := 1; j <= len(x[i].parts[j]); j++ {
			fmt.Printf("│   parts: %-22s│\n", x[i].parts[j])
		}
		fmt.Printf("│   Total: %-22v│", x[i].total)
		fmt.Println()
		k++
	}
	fmt.Println("╘════════════════════════════════╛")
}

// I.S. terdefinisi data array x yang berisi sebuah data pelanggan
// F.S. menampilkan hasil selection sort
func menurunBulan(x *arrPengguna, n *int) {
	for i := 0; i < *n; i++ {
		maxIdx := i
		for j := i + 1; j < *n; j++ {
			if x[j].bulanService > x[maxIdx].bulanService {
				maxIdx = j
			}
		}
		x[i], x[maxIdx] = x[maxIdx], x[i]
	}
	k := 1
	fmt.Println("╒════════════════════════════════╕")
	fmt.Println("│   Data History Bulan Terlama   │")
	fmt.Println("├────────────────────────────────┤")
	for i := 0; i < *n; i++ {
		fmt.Printf("│%v. Nama : %-22s│\n", k, x[i].nama)
		fmt.Printf("│   Bulan: %-22v│\n", x[i].bulanService)
		for j := 1; j <= len(x[i].parts[j]); j++ {
			fmt.Printf("│   parts: %-22s│\n", x[i].parts[j])
		}
		fmt.Printf("│   Total: %-22v│", x[i].total)
		fmt.Println()
		k++
	}
	fmt.Println("╘════════════════════════════════╛")
}

// I.S. terdefinisi data array y yang berisi sebuah data spare parts
// F.S. menampilkan hasil selection sort
func SpareTerlaris(y *arrParts, m int) {
	i := 0
	maxIdx := i
	for j := i + 1; j < m; j++ {
		if y[j].terjual > y[maxIdx].terjual {
			maxIdx = j
		}
	}
	y[i], y[maxIdx] = y[maxIdx], y[i]
	k := 1
	fmt.Println("Data Parts : ")
	fmt.Println("╒════════════════════════════════╕")
	fmt.Println("│          Data Terlaris         │")
	fmt.Println("├────────────────────────────────┤")
	for i := 0; i < m; i++ {
		fmt.Printf("│%v. parts: %-22s│\n", k, y[i].spare)
		fmt.Printf("│   Total: %-22v│", y[i].terjual)
		fmt.Println()
		k++
	}
	fmt.Println("╘════════════════════════════════╛")
}

// data bulan terbaru muncul
// input nama dan bulan
// ouput nama,bulan,spare,total

/*
func tampilData(hasilRoti, hasilMakaroni, hasilPizza int) {
	fmt.Println("Jumlah makanan:")
	fmt.Println("Roti:", hasilRoti)
	fmt.Println("Makaroni:", hasilMakaroni)
	fmt.Println("Pizza:", hasilPizza)

	// Tampilkan tabel hasil makanan
	fmt.Println("\nTabel Hasil Makanan:")
	fmt.Println("╒═══════════╤═════╤═════════╕")
	fmt.Println("│ Nama      │ Age │ Job     │")
	fmt.Println("├───────────┼─────┼─────────┤")
	fmt.Printf("│ %-9s │ %-3d │ %-7d │\n", "Roti", hasilRoti, 0)
	fmt.Printf("│ %-9s │ %-3d │ %-7d │\n", "Makaroni", hasilMakaroni, 0)
	fmt.Printf("│ %-9s │ %-3d │ %-7d │\n", "Pizza", hasilPizza, 0)
	fmt.Println("╘═══════════╧═════╧═════════╛")
}
*/

/*
func tampilData(x arrPengguna, n *int) {
	k := 1
	fmt.Println("╒═══════════════════════════════════════════╕")
	fmt.Println("│Data History :                             │")
	fmt.Println("├───────────────────────────────────────────┤")
	for i := 0; i < *n; i++ {
		fmt.Printf("│%v. %-23s                 │\n", k, x[i].nama)
		fmt.Printf("│   Bulan Service:  %-22v  │\n", x[i].bulanService)
		for j := 1; j <= len(x[i].parts[j]); j++ {
			fmt.Printf("│   %-22s                  │\n", x[i].parts[j])
		}
		fmt.Printf("│   %-22v                  │", x[i].total)
		fmt.Println()
		k++
	}
	fmt.Println("╘═══════════════════════════════════════════╛")
}
*/

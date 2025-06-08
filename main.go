package main

import (
	"fmt"
	"strings"
)

type Tanggal struct {
	Hari, Bulan, Tahun int
}

type Kendaraan struct {
	Plat     string
	Nama     string
	Jenis    string
	Harga    int
	Tersedia bool
}

type Penyewaan struct {
	IDSewa       string
	PlatKendaraan string
	NamaPenyewa  string
	TglSewa      Tanggal
	LamaSewa     int
	TotalBayar   int
}

const MaksData = 100

var daftarKendaraan [MaksData]Kendaraan
var daftarSewa [MaksData]Penyewaan
var jumlahKendaraan, jumlahSewa int
var totalPendapatan int

func main() {
	menu()
}

func menu() {
	pilihan := 0
	for pilihan != 8 {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Input Data Kendaraan")
		fmt.Println("2. Tampilkan Data Kendaraan")
		fmt.Println("3. Sewa Kendaraan")
		fmt.Println("4. Edit Data Kendaraan")
		fmt.Println("5. Urutkan Kendaraan")
		fmt.Println("6. Pengembalian Kendaraan")
		fmt.Println("7. Lihat Total Pendapatan")
		fmt.Println("8. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			inputKendaraan()
		case 2:
			tampilkanKendaraan()
		case 3:
			sewaKendaraan()
		case 4:
			editDataKendaraan()
		case 5:
			menuUrutkan()
		case 6:
			pengembalianKendaraan()
		case 7:
			fmt.Printf("Total Pendapatan: Rp%d\n", totalPendapatan)
		case 8:
			fmt.Println("Terima kasih!")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func inputKendaraan() {
	if jumlahKendaraan >= MaksData {
		fmt.Println("Data kendaraan penuh.")
		return
	}
	fmt.Println("\nInput Data Kendaraan")
	fmt.Print("Plat: ")
	fmt.Scan(&daftarKendaraan[jumlahKendaraan].Plat)
	fmt.Print("Nama: ")
	fmt.Scan(&daftarKendaraan[jumlahKendaraan].Nama)
	fmt.Print("Jenis: ")
	fmt.Scan(&daftarKendaraan[jumlahKendaraan].Jenis)
	fmt.Print("Harga per hari: ")
	fmt.Scan(&daftarKendaraan[jumlahKendaraan].Harga)
	daftarKendaraan[jumlahKendaraan].Tersedia = true
	jumlahKendaraan++
}

func tampilkanKendaraan() {
	fmt.Println("\nDaftar Kendaraan:")
	for i := 0; i < jumlahKendaraan; i++ {
		k := daftarKendaraan[i]
		fmt.Printf("%d. [%s] %s - %s - Rp%d/hari - Tersedia: %v\n", i+1, k.Plat, k.Nama, k.Jenis, k.Harga, k.Tersedia)
	}
}

func sewaKendaraan() {
	var plat string
	fmt.Println("\nPenyewaan Kendaraan")
	fmt.Print("Masukkan Plat Kendaraan: ")
	fmt.Scan(&plat)
	idx := cariKendaraanSequential(plat)
	if idx == -1 || !daftarKendaraan[idx].Tersedia {
		fmt.Println("Kendaraan tidak ditemukan atau tidak tersedia.")
		return
	}
	sewa := Penyewaan{}
	fmt.Print("ID Sewa: ")
	fmt.Scan(&sewa.IDSewa)
	sewa.PlatKendaraan = plat
	fmt.Print("Nama Penyewa: ")
	fmt.Scan(&sewa.NamaPenyewa)
	fmt.Print("Tanggal Sewa (dd mm yyyy): ")
	fmt.Scan(&sewa.TglSewa.Hari, &sewa.TglSewa.Bulan, &sewa.TglSewa.Tahun)
	fmt.Print("Lama Sewa (hari): ")
	fmt.Scan(&sewa.LamaSewa)
	sewa.TotalBayar = daftarKendaraan[idx].Harga * sewa.LamaSewa
	daftarSewa[jumlahSewa] = sewa
	jumlahSewa++
	daftarKendaraan[idx].Tersedia = false
	fmt.Println("Sewa berhasil dicatat. Total bayar: Rp", sewa.TotalBayar)
}

func pengembalianKendaraan() {
	var id string
	fmt.Print("Masukkan ID Sewa untuk pengembalian: ")
	fmt.Scan(&id)
	for i := 0; i < jumlahSewa; i++ {
		if daftarSewa[i].IDSewa == id {
			platKend := daftarSewa[i].PlatKendaraan
			idxK := cariKendaraanSequential(platKend)
			if idxK != -1 {
				daftarKendaraan[idxK].Tersedia = true
				totalPendapatan += daftarSewa[i].TotalBayar
				fmt.Printf("Kendaraan %s telah dikembalikan.\n", platKend)
				return
			}
		}
	}
	fmt.Println("ID Sewa tidak ditemukan.")
}

func cariKendaraanSequential(plat string) int {
	for i := 0; i < jumlahKendaraan; i++ {
		if strings.EqualFold(daftarKendaraan[i].Plat, plat) {
			return i
		}
	}
	return -1
}

func editDataKendaraan() {
	var plat string
	fmt.Print("Masukkan Plat Kendaraan yang ingin diedit: ")
	fmt.Scan(&plat)
	idx := cariKendaraanSequential(plat)
	if idx == -1 {
		fmt.Println("Kendaraan tidak ditemukan.")
		return
	}
	fmt.Println("Data lama akan diganti.")
	fmt.Print("Nama baru: ")
	fmt.Scan(&daftarKendaraan[idx].Nama)
	fmt.Print("Jenis baru: ")
	fmt.Scan(&daftarKendaraan[idx].Jenis)
	fmt.Print("Harga baru: ")
	fmt.Scan(&daftarKendaraan[idx].Harga)
}

// Quick Sort untuk Harga
func quickSortHarga(arr []Kendaraan, low, high int, asc bool) {
	if low < high {
		pi := partitionHarga(arr, low, high, asc)
		quickSortHarga(arr, low, pi-1, asc)
		quickSortHarga(arr, pi+1, high, asc)
	}
}

func partitionHarga(arr []Kendaraan, low, high int, asc bool) int {
	pivot := arr[high].Harga
	i := low - 1
	for j := low; j < high; j++ {
		if (asc && arr[j].Harga < pivot) || (!asc && arr[j].Harga > pivot) {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Heap Sort untuk Nama
func heapifyNama(arr []Kendaraan, n, i int, asc bool) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	shouldSwap := func(a, b string) bool {
		if asc {
			return strings.ToLower(a) > strings.ToLower(b)
		}
		return strings.ToLower(a) < strings.ToLower(b)
	}

	if l < n && shouldSwap(arr[largest].Nama, arr[l].Nama) {
		largest = l
	}
	if r < n && shouldSwap(arr[largest].Nama, arr[r].Nama) {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapifyNama(arr, n, largest, asc)
	}
}

func heapSortNama(arr []Kendaraan, asc bool) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapifyNama(arr, n, i, asc)
	}
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapifyNama(arr, i, 0, asc)
	}
}

// Menu Urutkan dengan QuickSort & HeapSort
func menuUrutkan() {
	var kategori, urutan int
	fmt.Println("\nKategori Urut:")
	fmt.Println("1. Berdasarkan Harga")
	fmt.Println("2. Berdasarkan Nama")
	fmt.Print("Pilihan: ")
	fmt.Scan(&kategori)
	fmt.Print("Urutan (1=Asc, 2=Desc): ")
	fmt.Scan(&urutan)
	asc := urutan == 1

	if kategori == 1 {
		quickSortHarga(daftarKendaraan[:jumlahKendaraan], 0, jumlahKendaraan-1, asc)
	} else {
		heapSortNama(daftarKendaraan[:jumlahKendaraan], asc)
	}

	tampilkanKendaraan()
}

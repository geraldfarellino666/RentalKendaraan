package main

import (
	"fmt"
	"strings"
)

// === Struct ===
type Kendaraan struct {
	ID     int
	Nama   string
	Tipe   string
	Status bool // true = tersedia, false = disewa
}

type Pelanggan struct {
	ID   int
	Nama string
}

type Rental struct {
	IDKendaraan int
	IDPelanggan int
	Hari        int
}

// === Array data ===
var dataKendaraan = []Kendaraan{
	{1, "Avanza", "Mobil", true},
	{2, "NMax", "Motor", true},
	{3, "Brio", "Mobil", true},
}

var dataPelanggan = []Pelanggan{
	{1, "Andi"},
	{2, "Budi"},
	{3, "Citra"},
}

var dataRental []Rental

// === Fungsi Searching ===
// Linear Search untuk mencari kendaraan berdasarkan nama
func cariKendaraanNama(nama string) *Kendaraan {
	for i := 0; i < len(dataKendaraan); i++ {
		if strings.EqualFold(dataKendaraan[i].Nama, nama) {
			return &dataKendaraan[i]
		}
	}
	return nil
}

// Binary Search (harus sudah terurut berdasarkan ID)
func binarySearchKendaraanByID(arr []Kendaraan, id int) *Kendaraan {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid].ID == id {
			return &arr[mid]
		} else if arr[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

// === Fungsi Sorting ===
// Quick Sort berdasarkan nama kendaraan
func quickSortKendaraanByNama(arr []Kendaraan, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSortKendaraanByNama(arr, low, p-1)
		quickSortKendaraanByNama(arr, p+1, high)
	}
}

func partition(arr []Kendaraan, low, high int) int {
	pivot := arr[high].Nama
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j].Nama < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Heap Sort berdasarkan ID kendaraan
func heapSort(arr []Kendaraan) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []Kendaraan, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left].ID > arr[largest].ID {
		largest = left
	}
	if right < n && arr[right].ID > arr[largest].ID {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// === Fungsi Rental ===
func sewaKendaraan(idKendaraan, idPelanggan, hari int) {
	kendaraan := binarySearchKendaraanByID(dataKendaraan, idKendaraan)
	if kendaraan == nil {
		fmt.Println("Kendaraan tidak ditemukan.")
		return
	}
	if !kendaraan.Status {
		fmt.Println("Kendaraan sedang disewa.")
		return
	}
	kendaraan.Status = false
	dataRental = append(dataRental, Rental{idKendaraan, idPelanggan, hari})
	fmt.Println("Berhasil menyewa kendaraan.")
}

// === Main Program ===
func main() {
	fmt.Println("=== Aplikasi Rental Kendaraan ===")

	fmt.Println("\nData Kendaraan:")
	for _, k := range dataKendaraan {
		fmt.Printf("%d. %s (%s) - Tersedia: %t\n", k.ID, k.Nama, k.Tipe, k.Status)
	}

	fmt.Println("\nMelakukan penyortiran kendaraan berdasarkan nama (Quick Sort)...")
	quickSortKendaraanByNama(dataKendaraan, 0, len(dataKendaraan)-1)
	for _, k := range dataKendaraan {
		fmt.Printf("%d. %s\n", k.ID, k.Nama)
	}

	fmt.Println("\nMelakukan penyortiran kendaraan berdasarkan ID (Heap Sort)...")
	heapSort(dataKendaraan)
	for _, k := range dataKendaraan {
		fmt.Printf("%d. %s\n", k.ID, k.Nama)
	}

	fmt.Println("\nMencari kendaraan bernama 'Brio' (Linear Search)...")
	found := cariKendaraanNama("Brio")
	if found != nil {
		fmt.Printf("Ditemukan: %s (%s)\n", found.Nama, found.Tipe)
	}

	fmt.Println("\nMenyewa kendaraan ID 1 oleh pelanggan ID 2 selama 3 hari...")
	sewaKendaraan(1, 2, 3)

	fmt.Println("\nData Rental:")
	for _, r := range dataRental {
		fmt.Printf("Kendaraan ID %d disewa oleh Pelanggan ID %d selama %d hari\n", r.IDKendaraan, r.IDPelanggan, r.Hari)
	}
}

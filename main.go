package main

import (
	"fmt"
	"sort"
	"strings"
)

// Struct
type Vehicle struct {
	Plate        string
	Name         string
	Type         string
	IsRented     bool
	HargaPerHari int
}

type Customer struct {
	ID   int
	Name string
}

type Rental struct {
	RentalID     int
	CustomerID   int
	Plate        string
	LamaHari     int
	TotalBiaya   int
}

var vehicles []Vehicle
var customers []Customer
var rentals []Rental

var nextCustomerID int = 1
var nextRentalID int = 1

func main() {
	for {
		fmt.Println("\n=== Aplikasi Rental Kendaraan ===")
		fmt.Println("1. Tambah Kendaraan")
		fmt.Println("2. Tambah Customer")
		fmt.Println("3. Tampilkan Kendaraan")
		fmt.Println("4. Tampilkan Customer")
		fmt.Println("5. Rental Kendaraan")
		fmt.Println("6. Kembalikan Kendaraan")
		fmt.Println("7. Cari Kendaraan (nama)")
		fmt.Println("8. Cari Customer (nama)")
		fmt.Println("9. Urutkan Kendaraan (nama)")
		fmt.Println("10. Urutkan Customer (nama)")
		fmt.Println("11. Lihat Pendapatan Total")
		fmt.Println("12. Keluar")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addVehicle()
		case 2:
			addCustomer()
		case 3:
			listVehicles()
		case 4:
			listCustomers()
		case 5:
			rentVehicle()
		case 6:
			returnVehicle()
		case 7:
			searchVehicleByName()
		case 8:
			searchCustomerByName()
		case 9:
			sortVehiclesByName()
		case 10:
			sortCustomersByName()
		case 11:
			showTotalRevenue()
		case 12:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func addVehicle() {
	var name, vehicleType, plate string
	var harga int
	fmt.Print("Masukkan plat kendaraan: ")
	fmt.Scanln(&plate)
	fmt.Print("Masukkan nama kendaraan: ")
	fmt.Scanln(&name)
	fmt.Print("Masukkan tipe kendaraan (Mobil/Motor): ")
	fmt.Scanln(&vehicleType)
	fmt.Print("Masukkan harga sewa per hari: ")
	fmt.Scanln(&harga)

	vehicle := Vehicle{
		Plate:        plate,
		Name:         name,
		Type:         vehicleType,
		HargaPerHari: harga,
		IsRented:     false,
	}
	vehicles = append(vehicles, vehicle)
	fmt.Println("Kendaraan berhasil ditambahkan.")
}

func addCustomer() {
	var name string
	fmt.Print("Masukkan nama customer: ")
	fmt.Scanln(&name)

	customer := Customer{
		ID:   nextCustomerID,
		Name: name,
	}
	customers = append(customers, customer)
	nextCustomerID++
	fmt.Println("Customer berhasil ditambahkan.")
}

func listVehicles() {
	fmt.Println("\nDaftar Kendaraan:")
	if len(vehicles) == 0 {
		fmt.Println("Belum ada kendaraan.")
		return
	}
	for _, v := range vehicles {
		status := "Tersedia"
		if v.IsRented {
			status = "Dirental"
		}
		fmt.Printf("Plat: %s | Nama: %s | Tipe: %s | Harga/hari: %d | Status: %s\n",
			v.Plate, v.Name, v.Type, v.HargaPerHari, status)
	}
}

func listCustomers() {
	fmt.Println("\nDaftar Customer:")
	if len(customers) == 0 {
		fmt.Println("Belum ada customer.")
		return
	}
	for _, c := range customers {
		fmt.Printf("ID: %d | Nama: %s\n", c.ID, c.Name)
	}
}

func rentVehicle() {
	var plate string
	var customerID, hari int
	fmt.Print("Masukkan plat kendaraan: ")
	fmt.Scanln(&plate)
	fmt.Print("Masukkan ID customer: ")
	fmt.Scanln(&customerID)
	fmt.Print("Berapa hari ingin dirental? ")
	fmt.Scanln(&hari)

	for i, v := range vehicles {
		if v.Plate == plate {
			if v.IsRented {
				fmt.Println("Kendaraan sedang dirental.")
				return
			}
			vehicles[i].IsRented = true

			total := v.HargaPerHari * hari
			rental := Rental{
				RentalID:   nextRentalID,
				CustomerID: customerID,
				Plate:      plate,
				LamaHari:   hari,
				TotalBiaya: total,
			}
			rentals = append(rentals, rental)
			nextRentalID++
			fmt.Printf("Kendaraan berhasil dirental. Total biaya: %d\n", total)
			return
		}
	}
	fmt.Println("Kendaraan tidak ditemukan.")
}

func returnVehicle() {
	var plate string
	fmt.Print("Masukkan plat kendaraan yang ingin dikembalikan: ")
	fmt.Scanln(&plate)

	for i, v := range vehicles {
		if v.Plate == plate {
			if !v.IsRented {
				fmt.Println("Kendaraan belum dirental.")
				return
			}
			vehicles[i].IsRented = false
			fmt.Println("Kendaraan berhasil dikembalikan.")
			return
		}
	}
	fmt.Println("Kendaraan tidak ditemukan.")
}

func searchVehicleByName() {
	var keyword string
	fmt.Print("Masukkan nama kendaraan yang dicari: ")
	fmt.Scanln(&keyword)

	found := false
	for _, v := range vehicles {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(keyword)) {
			status := "Tersedia"
			if v.IsRented {
				status = "Dirental"
			}
			fmt.Printf("Plat: %s | Nama: %s | Tipe: %s | Harga/hari: %d | Status: %s\n",
				v.Plate, v.Name, v.Type, v.HargaPerHari, status)
			found = true
		}
	}
	if !found {
		fmt.Println("Kendaraan tidak ditemukan.")
	}
}

func searchCustomerByName() {
	var keyword string
	fmt.Print("Masukkan nama customer yang dicari: ")
	fmt.Scanln(&keyword)

	found := false
	for _, c := range customers {
		if strings.Contains(strings.ToLower(c.Name), strings.ToLower(keyword)) {
			fmt.Printf("ID: %d | Nama: %s\n", c.ID, c.Name)
			found = true
		}
	}
	if !found {
		fmt.Println("Customer tidak ditemukan.")
	}
}

func sortVehiclesByName() {
	sort.Slice(vehicles, func(i, j int) bool {
		return strings.ToLower(vehicles[i].Name) < strings.ToLower(vehicles[j].Name)
	})
	fmt.Println("Kendaraan berhasil diurutkan berdasarkan nama.")
	listVehicles()
}

func sortCustomersByName() {
	sort.Slice(customers, func(i, j int) bool {
		return strings.ToLower(customers[i].Name) < strings.ToLower(customers[j].Name)
	})
	fmt.Println("Customer berhasil diurutkan berdasarkan nama.")
	listCustomers()
}

func showTotalRevenue() {
	total := 0
	for _, r := range rentals {
		total += r.TotalBiaya
	}
	fmt.Printf("Total pendapatan dari rental: %d\n", total)
}

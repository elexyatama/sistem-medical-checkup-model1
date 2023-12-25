package main

import (
	"fmt"
)

const nmax_patient int = 100
const nmax_pack int = 50
const nmax_mcu int = 500

type info_patient struct {
	name   string
	id     string
	origin string
	age    int
	gender string
}

type info_pack struct {
	name     string
	id       string
	category string
}

type info_mcu struct { //mcu = medical check up
	id      string
	price   float64
	period  string //periode, waktu mcu
	patient info_patient
	pack    info_pack
}

type patient_tab struct {
	data [nmax_patient]info_patient
	n    int //buat ngitung ada brp data dalam list/array
}

type pack_tab struct {
	data [nmax_pack]info_pack
	n    int //buat ngitung ada brp data dalam list/array
}

type mcu_tab struct {
	data [nmax_patient]info_mcu
	n    int //buat ngitung ada brp data dalam list/array
}

func main() {
	var patients patient_tab
	var packs pack_tab
	var mcues mcu_tab
	var input int = 99
	for input != 7 {
		fmt.Println("------------ Medical Data ------------------")
		print_cmd_main()
		fmt.Scan(&input)
		if input == 1 {
			view_mcues(mcues)
		} else if input == 2 {
			view_patients(patients)
		} else if input == 3 {
			view_packs(packs)
		} else if input == 4 {
			register_mcu(&mcues, packs, patients)
		} else if input == 5 {
			finish_mcu(&mcues)
		} else if input == 6 {
			search_menu(mcues, patients)
		} else if input == 7 {
			edit_menu(&packs, &patients)
		} else if input == 0 {
			fmt.Println("Exiting program")
		} else {
			fmt.Println("Invalid input. Please enter a number between 0 and 7.")
			fmt.Scanln()
		}
	}
}

func view_mcues(mcues mcu_tab) {
	var input int = 99
	for input != 0 {
		fmt.Println("----------------- MCU ----------------------")
		print_all_mcu(mcues)
		fmt.Scan(&input)
		if input == 1 {
			sort_period(&mcues)
		} else if input == 2 {
			sort_pack(&mcues)
		} else if input == 0 {
			fmt.Println("Returning to main menu")
		} else {
			fmt.Println("Unknown input")
			fmt.Scanln()
		}
	}
}

func view_patients(patients patient_tab) {
	fmt.Println("--------------- Patient --------------------")
	print_all_patient(patients)
	fmt.Print("Press anywhere to return")
	fmt.Scanln()
}

func view_packs(packs pack_tab) {
	fmt.Println("---------------- Packs ---------------------")
	print_all_packs(packs)
	fmt.Print("Press anywhere to return")
	fmt.Scanln()
}

func add_patient(patients *patient_tab, x info_patient) { //buat nambahin data pasien
	patients.data[patients.n] = x //masukin nilai x, yang nanti berupa input dari func main
	patients.n++                  //tiap func add ini dipanggil, variabel n harus ditambah buat track brp banyk data dalam array
}

func add_pack(packs *pack_tab, x info_pack) { // sistem nya sama kae add_patient, cuma ini buat pack
	packs.data[packs.n] = x
	packs.n++
}

func add_mcu(mcues *mcu_tab, x info_mcu) { // sistem nya sama kae add_patient, cuma ini buat mcu
	mcues.data[mcues.n] = x
	mcues.n++
}

func remove_patient(patients *patient_tab, x int) {
	for i := x; i < patients.n-1; i++ { /* logika nya ini misal kita ada array 1 2 3 4 5, kalo mau hapus angka 3 */
		patients.data[i] = patients.data[i+1] /*  harus geser stu satu sampe ujung, biar si angka 3 itu ke tumpuk dan nilainya ilang*/
	} /*  awalnya gni 1 2 3 4 5, trus nanti gr2 geser jadi 1 2 4 5 0*/
	patients.n-- /*  angka kanannya geser trus ke kiri mulai dari angka ke 3*/
}

func remove_pack(packs *pack_tab, x int) { //konsepnya sama kae remove_patient
	for i := x; i < packs.n-1; i++ {
		packs.data[i] = packs.data[i+1]
	}
	packs.n--
}

func remove_mcu(mcues *mcu_tab, x int) { //konsepnya sama kae remove_patient
	for i := x; i < mcues.n-1; i++ {
		mcues.data[i] = mcues.data[i+1]
	}
	mcues.n--
}

func search_patient_from_pack(mcues mcu_tab, x string) { // ini buat nampilin semua pasien dg pack tertentu, x itu parameter pack yang ingin diprint semua pasien terkaitnya
	var count int = 1
	for i := 0; i < mcues.n; i++ { // ini ngecek dari array index 0 sampe mcues.n, mcues.n ini tu ukuran array
		if mcues.data[i].pack.id == x { // ini ngecek kalo misalnya data.pack.id sama kae apa yang kita cari, nanti diprint detailnya dibawah
			print_mcu_detail(mcues.data[i])
			count++
		}
	}
	if count == 1 { //ini kondisi kalo ngga ada pasien dg pack yg pengen kita cari, indikasinya count bernilai 1
		fmt.Print("No patient with that spesific pack listed")
	}
}

func search_patient_from_period(mcues mcu_tab, x string) { //sequential search ada disini
	var count int = 1
	for i := 0; i < mcues.n; i++ {
		if mcues.data[i].period == x {
			print_mcu_detail(mcues.data[i])
			count++
		}
	}
	if count == 1 {
		fmt.Print("No patient with that spesific period listed")
	}
}

func search_patient_name(patients patient_tab, x string) int { //binary search ada disini
	sort_patient_name(&patients)
	low, high := 0, patients.n-1
	found := -1

	for low <= high {
		mid := (low + high) / 2
		curr := patients.data[mid].name

		if curr == x {
			found = mid
			high = mid - 1
		} else if curr < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found != -1 { //found kalo ngga -1 artinya data nya ditemuin, kalo ditemuin nanti langsung print detailnya
		print_patient_detail(patients.data[found])
	} else { //klo ga nemu pass error msg
		fmt.Printf("Patient with ID %s not found\n", x)
	}
	return found
}

func search_pack_id(packs pack_tab, x string) int { //sequential search disini
	for i := 0; i < packs.n; i++ {
		if packs.data[i].id == x {
			return i
		}
	}
	return -1
}

func search_patient_id(patients patient_tab, x string) int { //sequential search disini
	for i := 0; i < patients.n; i++ {
		if patients.data[i].id == x {
			return i
		}
	}
	return -1
}

func search_mcu_id(mcues mcu_tab, x string) int { //sequential search disini
	for i := 0; i < mcues.n; i++ {
		if mcues.data[i].id == x {
			return i
		}
	}
	return -1
}

func sort_period(mcues *mcu_tab) { //selection sort ada disini
	for i := 0; i < mcues.n-1; i++ {
		minIndex := i
		for j := i + 1; j < mcues.n; j++ {
			if mcues.data[j].period < mcues.data[minIndex].period {
				minIndex = j
			}
		}
		mcues.data[i], mcues.data[minIndex] = mcues.data[minIndex], mcues.data[i]
	}
}

func sort_pack(mcues *mcu_tab) { //selection sort ada disini
	for i := 0; i < mcues.n-1; i++ {
		minIndex := i
		for j := i + 1; j < mcues.n; j++ {
			if mcues.data[j].pack.name < mcues.data[minIndex].pack.name {
				minIndex = j
			}
		}
		mcues.data[i], mcues.data[minIndex] = mcues.data[minIndex], mcues.data[i]
	}
}

func sort_patient_name(patients *patient_tab) { //insertion sort ada disini
	for i := 1; i < patients.n; i++ {
		key := patients.data[i]
		j := i - 1

		for j >= 0 && patients.data[j].name > key.name {
			patients.data[j+1] = patients.data[j]
			j--
		}
		patients.data[j+1] = key
	}
}

// func2 dibawah ini algoritma biasa buat ngeprint, sesuai dengan judul func nya

func print_all_patient(patients patient_tab) { // buat print semua list pasien
	fmt.Println("Patient Data:")
	for i := 0; i < patients.n; i++ {
		fmt.Printf("Name: %s\n", patients.data[i].name)
		fmt.Printf("ID: %s\n", patients.data[i].id)
		fmt.Printf("Origin: %s\n", patients.data[i].origin)
		fmt.Printf("Age: %d\n", patients.data[i].age)
		fmt.Printf("Gender: %s\n", patients.data[i].gender)
		fmt.Println("--------------")
	}
}

func print_all_packs(packs pack_tab) { // buat prin semua list pack
	fmt.Println("Pack Data:")
	for i := 0; i < packs.n; i++ {
		fmt.Printf("Name: %s\n", packs.data[i].name)
		fmt.Printf("ID: %s\n", packs.data[i].id)
		fmt.Printf("Category: %s\n", packs.data[i].category)
		fmt.Println("--------------")
	}
}

func print_all_mcu(mcues mcu_tab) { // buat print semua list mcu
	fmt.Println("MCU Data:")
	for i := 0; i < mcues.n; i++ {
		fmt.Printf("Patient: %s\n", mcues.data[i].patient.name)
		fmt.Printf("ID: %s\n", mcues.data[i].id)
		fmt.Printf("Price: %.2f\n", mcues.data[i].price)
		fmt.Printf("Period: %s\n", mcues.data[i].period)
		fmt.Printf("Pack: %s\n", mcues.data[i].pack.name)
		fmt.Println("--------------")
	}
}

func print_patient_detail(x info_patient) { // buat print detail satu pasien
	fmt.Println("Patient Details:")
	fmt.Printf("Name: %s\n", x.name)
	fmt.Printf("ID: %s\n", x.id)
	fmt.Printf("Origin: %s\n", x.origin)
	fmt.Printf("Age: %d\n", x.age)
	fmt.Printf("Gender: %s\n", x.gender)
	fmt.Println("--------------")
}

func print_mcu_detail(mcu info_mcu) {
	fmt.Println("MCU Details:")
	fmt.Printf("MCU ID: %s\n", mcu.id)
	fmt.Printf("Price: %.2f\n", mcu.price)
	fmt.Printf("Period: %s\n", mcu.period)
	fmt.Printf("Patient Name: %s\n", mcu.patient.name)
	fmt.Printf("Pack Name: %s\n", mcu.pack.name)
	fmt.Println("--------------")
}

func print_cmd_main() {
	fmt.Println("1. View MCU-es")
	fmt.Println("2. View patients")
	fmt.Println("3. View packs")
	fmt.Println("4. Make medical check up")
	fmt.Println("5. Finish medical check up")
	fmt.Println("6. Search data")
	fmt.Println("7. Edit data")
	fmt.Println("0. Exit program")
	fmt.Println("--------------------------------------------")
	fmt.Print("Input : ")
}

func print_cmd_view() {
	fmt.Println("1. Sort by period")
	fmt.Println("2. Sort by packages")
	fmt.Println("0. Back")
	fmt.Println("--------------------------------------------")
	fmt.Print("Input : ")
}

func register_mcu(mcues *mcu_tab, packs pack_tab, patients patient_tab) {
	// Periksa apakah ada paket medis yang tersedia
	if packs.n == 0 {
		fmt.Println("No medical packs available for registration.")
		return
	}

	// Periksa apakah ada pasien yang tersedia
	if patients.n == 0 {
		fmt.Println("No patients available for registration.")
		return
	}

	// Tampilkan paket medis yang tersedia untuk dipilih
	fmt.Println("Available Medical Packs:")
	for i := 0; i < packs.n; i++ {
		fmt.Printf("%d. %s\n", i+1, packs.data[i].name)
	}

	// Minta pengguna untuk memilih paket
	var selectedPackIndex int
	fmt.Print("Select a pack (enter the corresponding number): ")
	fmt.Scan(&selectedPackIndex)

	// Validasi indeks paket yang dipilih
	if selectedPackIndex < 1 || selectedPackIndex > packs.n {
		fmt.Println("Invalid pack selection.")
		return
	}

	// Tampilkan pasien yang tersedia untuk dipilih
	fmt.Println("\nAvailable Patients:")
	for i := 0; i < patients.n; i++ {
		fmt.Printf("%d. %s\n", i+1, patients.data[i].name)
	}

	// select patient
	var selectedPatientIndex int
	fmt.Print("Select a patient (enter the corresponding number): ")
	fmt.Scan(&selectedPatientIndex)

	// Validasi indeks pasien yang dipilih
	if selectedPatientIndex < 1 || selectedPatientIndex > patients.n {
		fmt.Println("Invalid patient selection.")
		return
	}

	var newMCU info_mcu

	fmt.Print("Enter MCU ID: ")
	fmt.Scan(&newMCU.id)

	fmt.Print("Enter MCU Price: ")
	fmt.Scan(&newMCU.price)

	fmt.Print("Enter MCU Period: ")
	fmt.Scan(&newMCU.period)

	newMCU.patient = patients.data[selectedPatientIndex-1]
	newMCU.pack = packs.data[selectedPackIndex-1]

	add_mcu(mcues, newMCU)

	fmt.Printf("Medical Check-up successfully registered. MCU ID: %s\n", newMCU.id)
}

func finish_mcu(mcues *mcu_tab) {
	var mcuID string
	fmt.Print("Enter the MCU id: ")
	fmt.Scan(&mcuID)
	index := search_mcu_id(*mcues, mcuID)

	if index != -1 { // kalo ketemu mcu nya masuk ke if ini
		print_mcu_detail(mcues.data[index]) //print detail mcu sebelom di hapus

		remove_mcu(mcues, index) //buat hapus mcu

		fmt.Println("MCU finished and removed.")
	} else {
		fmt.Println("MCU with ID", mcuID, "not found.")
	}
}

func search_menu(mcues mcu_tab, patients patient_tab) {
	var input int = 99
	for input != 0 {
		fmt.Println("Search Menu:")
		fmt.Println("1. Search patients by pack")
		fmt.Println("2. Search patients by period")
		fmt.Println("3. Search patient by ID")
		fmt.Println("0. Return")
		fmt.Println("--------------------------------------------")

		fmt.Print("Input: ")
		fmt.Scan(&input)

		if input == 0 {
			fmt.Println("Exiting search menu.")
			break
		} else if input == 1 {
			var packID string
			fmt.Print("Enter pack ID: ")
			fmt.Scan(&packID)
			search_patient_from_pack(mcues, packID)
		} else if input == 2 {
			var period string
			fmt.Print("Enter period: ")
			fmt.Scan(&period)
			search_patient_from_period(mcues, period)
		} else if input == 3 {
			var patientID string
			fmt.Print("Enter patient ID: ")
			fmt.Scan(&patientID)
			search_patient_name(patients, patientID)
		} else {
			fmt.Println("Invalid input. Please enter a number between 0 and 3.")
		}
	}
}

func edit_menu(packs *pack_tab, patients *patient_tab) {
	var input int = 99
	for input != 0 {
		fmt.Println("Edit Menu:")
		fmt.Println("1. Add patient")
		fmt.Println("2. Add pack")
		fmt.Println("3. Remove patient")
		fmt.Println("4. Remove pack")
		fmt.Println("5. Edit patient")
		fmt.Println("6. Edit pack")
		fmt.Println("0. Return")
		fmt.Println("--------------------------------------------")
		fmt.Print("input: ")
		fmt.Scan(&input)

		if input == 0 {
			fmt.Println("Exiting edit menu.")
		} else if input == 1 {
			var newPatient info_patient
			newPatient = get_input_patient()
			add_patient(&*patients, newPatient)
		} else if input == 2 {
			var newPack info_pack
			newPack = get_input_pack()
			add_pack(&*packs, newPack)
		} else if input == 3 {
			var patientID string
			fmt.Print("Enter patient ID to remove: ")
			fmt.Scan(&patientID)
			index := search_patient_id(*patients, patientID)
			if index != -1 {
				remove_patient(patients, index)
			} else {
				fmt.Println("Patient not found.")
			}
		} else if input == 4 {
			var packID string
			fmt.Print("Enter pack ID to remove: ")
			fmt.Scan(&packID)
			index := search_pack_id(*packs, packID)
			if index != -1 {
				remove_pack(packs, index)
			} else {
				fmt.Println("Pack not found.")
			}
		} else if input == 5 {
			var patientID string
			fmt.Print("Enter patient ID to edit: ")
			fmt.Scan(&patientID)
			index := search_patient_id(*patients, patientID)
			if index != -1 {
				fmt.Printf("Editing %s\n", patients.data[index].name)
				var newPatient info_patient
				newPatient = get_input_patient()
				patients.data[index] = newPatient
			} else {
				fmt.Println("Patient not found.")
			}
		} else if input == 6 {
			var packID string
			fmt.Print("Enter pack ID to edit: ")
			fmt.Scan(&packID)
			index := search_pack_id(*packs, packID)
			if index != -1 {
				fmt.Printf("Editing %s\n", packs.data[index].name)
				var newPack info_pack
				newPack = get_input_pack()
				packs.data[index] = newPack
			} else {
				fmt.Println("Pack not found.")
			}
		} else {
			fmt.Println("Invalid input. Please enter a number between 0 and 6.")
		}
	}
}

func get_input_patient() info_patient {
	var newPatient info_patient

	fmt.Print("Enter patient name: ")
	fmt.Scan(&newPatient.name)

	fmt.Print("Enter patient ID: ")
	fmt.Scan(&newPatient.id)

	fmt.Print("Enter patient origin: ")
	fmt.Scan(&newPatient.origin)

	fmt.Print("Enter patient age: ")
	fmt.Scan(&newPatient.age)

	fmt.Print("Enter patient gender: ")
	fmt.Scan(&newPatient.gender)

	return newPatient
}

func get_input_pack() info_pack {
	var newPack info_pack

	fmt.Print("Enter pack name: ")
	fmt.Scan(&newPack.name)

	fmt.Print("Enter pack ID: ")
	fmt.Scan(&newPack.id)

	fmt.Print("Enter pack category: ")
	fmt.Scan(&newPack.category)

	return newPack
}

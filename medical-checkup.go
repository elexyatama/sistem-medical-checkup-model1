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

type info_mcu struct {
	id      string
	price   float64
	period  string
	patient info_patient
	pack    info_pack
}

type patient_tab struct {
	data [nmax_patient]info_patient
	n    int
}

type pack_tab struct {
	data [nmax_pack]info_pack
	n    int
}

type mcu_tab struct {
	data [nmax_patient]info_mcu
	n    int
}

func main() {
	fmt.Print("yes")
}

func add_patient(patients *patient_tab, x info_patient) {
	patients.data[patients.n] = x
	patients.n++
}

func add_pack(packs *pack_tab, x info_pack) {
	packs.data[packs.n] = x
	packs.n++
}

func add_mcu(mcues *mcu_tab, x info_mcu) {
	mcues.data[mcues.n] = x
	mcues.n++
}

func remove_patient(patients *patient_tab, x int) {
	for i := x; i < patients.n-1; i++ {
		patients.data[i] = patients.data[i+1]
	}
	patients.n--
}

func remove_pack(packs *pack_tab, x int) {
	for i := x; i < packs.n-1; i++ {
		packs.data[i] = packs.data[i+1]
	}
	packs.n--
}

func remove_mcu(mcues *mcu_tab, x int) {
	for i := x; i < mcues.n-1; i++ {
		mcues.data[i] = mcues.data[i+1]
	}
	mcues.n--
}

func search_patient_from_pack(mcues mcu_tab, x string) {
	var count int = 1
	for i := 0; i < mcues.n; i++ {
		if mcues.data[i].pack.id == x {
			fmt.Printf("%d. %s, %s\n", count, mcues.data[i].patient.name, mcues.data[i].patient.id)
			count++
		}
	}
	if count == 1 {
		fmt.Print("No patient with that spesific pack listed")
	}
}

func search_patient_from_period(mcues mcu_tab, x string) { //sequential search
	var count int = 1
	for i := 0; i < mcues.n; i++ {
		if mcues.data[i].period == x {
			fmt.Printf("%d. %s, %s\n", count, mcues.data[i].patient.name, mcues.data[i].patient.id)
			count++
		}
	}
	if count == 1 {
		fmt.Print("No patient with that spesific period listed")
	}
}

func search_patient(patients patient_tab, x string) { //binary search
	sort_patient_id(&patients)
	low, high := 0, patients.n-1
	found := -1

	for low <= high {
		mid := (low + high) / 2
		curr := patients.data[mid].id

		if curr == x {
			found = mid
			high = mid - 1
		} else if curr < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found != -1 {
		print_patient_detail(patients.data[found])
	} else {
		fmt.Printf("Patient with ID %s not found\n", x)
	}
}

func sort_period(mcues *mcu_tab) { //selection sort
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

func sort_pack(mcues *mcu_tab) { //selection sort
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

func sort_patient_id(patients *patient_tab) { //insertion sort
	for i := 1; i < patients.n; i++ {
		key := patients.data[i]
		j := i - 1

		for j >= 0 && patients.data[j].id > key.id {
			patients.data[j+1] = patients.data[j]
			j--
		}
		patients.data[j+1] = key
	}
}

func print_all_patient(patients patient_tab) {
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

func print_all_pack(packs pack_tab) {
	fmt.Println("Pack Data:")
	for i := 0; i < packs.n; i++ {
		fmt.Printf("Name: %s\n", packs.data[i].name)
		fmt.Printf("ID: %s\n", packs.data[i].id)
		fmt.Printf("Category: %s\n", packs.data[i].category)
		fmt.Println("--------------")
	}
}

func print_all_mcu(mcues mcu_tab) {
	fmt.Println("MCU Data:")
	for i := 0; i < mcues.n; i++ {
		fmt.Printf("ID: %s\n", mcues.data[i].id)
		fmt.Printf("Price: %.2f\n", mcues.data[i].price)
		fmt.Printf("Period: %s\n", mcues.data[i].period)
		fmt.Printf("Patient: %s\n", mcues.data[i].patient.name)
		fmt.Printf("Pack: %s\n", mcues.data[i].pack.name)
		fmt.Println("--------------")
	}
}

func print_patient_detail(x info_patient) {
	fmt.Println("Patient Details:")
	fmt.Printf("Name: %s\n", x.name)
	fmt.Printf("ID: %s\n", x.id)
	fmt.Printf("Origin: %s\n", x.origin)
	fmt.Printf("Age: %d\n", x.age)
	fmt.Printf("Gender: %s\n", x.gender)
}

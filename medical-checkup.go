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

func add_patient(patient_list *patient_tab, x info_patient) {
	patient_list.data[patient_list.n] = x
	patient_list.n++
}

func add_pack(pack_list *pack_tab, x info_pack) {
	pack_list.data[pack_list.n] = x
	pack_list.n++
}

func add_mcu(mcu_list *mcu_tab, x info_mcu) {
	mcu_list.data[mcu_list.n] = x
	mcu_list.n++
}

func remove_patient(patient_list *patient_tab, x int) {
	for i := x; i < patient_list.n-1; i++ {
		patient_list.data[i] = patient_list.data[i+1]
	}
	patient_list.n--
}

func remove_pack(pack_list *pack_tab, x int) {
	for i := x; i < pack_list.n-1; i++ {
		pack_list.data[i] = pack_list.data[i+1]
	}
	pack_list.n--
}

func remove_mcu(mcu_list *mcu_tab, x int) {
	for i := x; i < mcu_list.n-1; i++ {
		mcu_list.data[i] = mcu_list.data[i+1]
	}
	mcu_list.n--
}

func search_patient_from_pack(mcu_list mcu_tab, x string) {
	var count int = 1
	for i := 0; i < mcu_list.n; i++ {
		if mcu_list.data[i].pack.id == x {
			fmt.Printf("%d. %s, %s\n", count, mcu_list.data[i].patient.name, mcu_list.data[i].patient.id)
			count++
		}
	}
	if count == 1 {
		fmt.Print("No patient with that spesific pack listed")
	}
}

func search_patient_from_period(mcu_list mcu_tab, x string) {
	var count int = 1
	for i := 0; i < mcu_list.n; i++ {
		if mcu_list.data[i].period == x {
			fmt.Printf("%d. %s, %s\n", count, mcu_list.data[i].patient.name, mcu_list.data[i].patient.id)
			count++
		}
	}
	if count == 1 {
		fmt.Print("No patient with that spesific period listed")
	}
}

func search_patient(patient_list patient_tab, x string) {
	for i := 0; i < patient_list.n; i++ {
		if patient_list.data[i].id == x {
			print_patient_detail(patient_list.data[i])
			return
		}
	}
	fmt.Print("No patient with that spesific id listed")
}

func sort_period(mcu_list *mcu_tab) {

}

func sort_pack(mcu_list *mcu_tab) {

}

func print_all_patient(patient_list patient_tab) {

}

func print_all_pack(pack_list pack_tab) {

}

func print_all_mcu(mcu_list mcu_tab) {

}

func print_patient_detail(x info_patient) {

}

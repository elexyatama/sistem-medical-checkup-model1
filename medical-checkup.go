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
	data [nmax_patient]info_patient
	n    int
}

func main() {
	fmt.Print("yes")
}

func add_patient(patient_list *patient_tab, x info_patient) {

}

func add_pack(pack_list *pack_tab, x info_pack) {

}

func add_mcu(mcu_list *mcu_tab, x info_mcu) {

}

func remove_patient(patient_list *patient_tab, x int) {

}

func remove_pack(pack_list *pack_tab, x int) {

}

func remove_mcu(mcu_list *mcu_tab, x int) {

}

func search_patient_from_pack(mcu_list mcu_tab, x string) {

}

func search_patient_from_period(msu_list mcu_tab, x string) {

}

func search_patient(patient_list patient_tab, x string) {

}

func sort_period(mcu_list *mcu_tab) {

}

func sort_pack(mcu_list *mcu_tab) {

}

func print_patient(patient_list patient_tab) {

}

func print_pack(pack_list pack_tab) {

}

func print_mcu(mcu_list mcu_tab) {

}

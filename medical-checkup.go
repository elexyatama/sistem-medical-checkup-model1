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

package main

import (
	"fmt"
)

// Variablen
var tage int
var kilometer int
var steigerung float32

func main() {
	// Bildschirm löschen
	fmt.Print("\033[H\033[2J")
	eingabe(kilometer)
}

func eingabe(kilometer int) {
	// Eingabe
	fmt.Print("Wieviele Killometer möchten Sie laufen ? ")
	fmt.Scan(&kilometer)
	fmt.Println("----------------------------------------")
	fmt.Println("")
	darstellung(kilometer)
}

func darstellung(kilometer int) {
	// Leistungssteigerung
	for i := 1; i <= kilometer*2; i++ {
		x := float32(i)
		x = x*0.5 + 0.5
		steigerung = 1 / x * 100
		if i%2 == 0 {
			fmt.Println("Tag:", i, "REGENERATION")
		} else {
			fmt.Print("Tag: ", i, " Kilometer: ", x, " Steigerung: ")
			fmt.Printf("%.1f", steigerung)
			fmt.Print("%")
			fmt.Println("")
		}
	}
	zusammenfassung(kilometer)
}

func zusammenfassung(kilometer int) {
	// Zusammenfassung
	tage = kilometer * 2
	fmt.Println("")
	fmt.Println("Für", kilometer, "Kilometer, benötigen Sie:", tage, "Tage")
	fmt.Println("")
}

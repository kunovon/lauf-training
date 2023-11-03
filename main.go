package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

// Variablen
var tage int
var kilometer int
var steigerung float32
var x float32

func main() {
	// Bildschirm löschen
	fmt.Print("\033[H\033[2J")

	nocheinmal := true

	for nocheinmal {
		eingabe(kilometer)
		nocheinmal = GetYesOrNo("Möchten Sie noch einmal? (j/n)?\n")

	}

	fmt.Println("")
	fmt.Println("Auf Wiedersehen!")
	fmt.Println("")
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
		x = float32(i)
		x = x*0.5 + 0.5
		steigerung = 1 / x * 100
		if i%2 == 0 {
			fmt.Printf("Tag: %d REGENERATION\n", i)
		} else {
			fmt.Printf("Tag: %d Kilometer: %.0f  Steigerung %.1f %%\n", i, x, steigerung)
		}
	}
	zusammenfassung(kilometer)
}

func zusammenfassung(kilometer int) {
	// Zusammenfassung
	tage = kilometer * 2
	fmt.Println("-----------------------------------------")
	fmt.Printf("Für %d Kilometer, benötigen Sie: %d Tage!\n", kilometer, tage)
	fmt.Printf("=========================================\n\n")
}

// GetYesOrNo erlaubt die erneute Eingabe
func GetYesOrNo(q string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(q)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'n' || char == 'N' {
			return false
		}
		return true
	}
}

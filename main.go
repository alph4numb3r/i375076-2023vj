package main

import (
	"fmt"
	"time"
)

// time constants
const (
        MorningThreshold = 7
        AfternoonThreshold = 12
        EveningThreshold = 18
        NightThreshold = 23
)

// message constants
const (
        DefaultMessage = "! Welkom bij Fonteyn Vakantieparken"
        MorningPrefix = "Goedemorgen"
        AfternoonPrefix = "Goedemiddag"
        EveningPrefix = "Goedenavond"
        MorningMessage = MorningPrefix + DefaultMessage
        AfternoonMessage = AfternoonPrefix + DefaultMessage
        EveningMessage = EveningPrefix + DefaultMessage
        NightMessage = "Sorry, de parkeerplaats is ’s nachts gesloten"
)

func main() {
    time := time.Now()

    switch {
    case time.Hour() < MorningThreshold:
        fmt.Println(NightMessage)
    case time.Hour() < AfternoonThreshold:
        fmt.Println(MorningMessage)
    case time.Hour() < EveningThreshold:
        fmt.Println(AfternoonMessage)
    case time.Hour() < NightThreshold:
        fmt.Println(EveningMessage)
    default:
        fmt.Println(NightMessage)
    }
}
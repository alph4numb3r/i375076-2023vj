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

// debug constants
/*
const (
        debugOffset = 6
)
*/

func main() {
    // debug
    /* 
    now := time.Now().Add(time.Hour * debugOffset)
    fmt.Println(now.Clock())
    */

    // production
    ///*
    now := time.Now()
    //*/

    switch {
    case now.Hour() < MorningThreshold:
        fmt.Println(NightMessage)
    case now.Hour() < AfternoonThreshold:
        fmt.Println(MorningMessage)
    case now.Hour() < EveningThreshold:
        fmt.Println(AfternoonMessage)
    case now.Hour() < NightThreshold:
        fmt.Println(EveningMessage)
    default:
        fmt.Println(NightMessage)
    }
}
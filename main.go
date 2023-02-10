package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	Set "github.com/deckarep/golang-set/v2"
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

//license plates
type TLicensePlates Set.Set[string]

// debug constants
const (
        //debugOffset = 6
)
func debugLicensePlates() []string {
    return []string{
        "1-ABC-123",
        "2-DEF-456",
        "3-GHI-789",
        "4-JKL-012",
        "5-MNO-345",
        "6-PQR-678",
        "7-STU-901",
        "8-VWX-234",
    }
}






func main() {
    var (
    // TEST: set time
    //now = time.Now().Add(time.Hour * debugOffset)
    // PROD: get time
    now = time.Now()

    reader = bufio.NewReader(os.Stdin)
    licensePlates TLicensePlates
    )


    // PROD: get license plates from database
    //licensePlates = 
    
    // TEST: hard-coded license plates
    licensePlates = Set.NewSet(debugLicensePlates()...)

    var licensePlate string
    // PROD: get license plate from camera

    // TEST: input license plate from console
    fmt.Print("License plate: ")
    for err := *new(error); err != nil; licensePlate, err = reader.ReadString('\n') {
        fmt.Print("License plate: ")
    }

    // check if license plate is in set
    if licensePlates.Contains(licensePlate) {
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
    } else {
        fmt.Println("U heeft helaas geen toegang tot het parkeerterrein")
    }
}
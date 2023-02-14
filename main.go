package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	Set "github.com/deckarep/golang-set/v2"
)

// time constants
const (
        MorningThreshold    = 7
        AfternoonThreshold  = 12
        EveningThreshold    = 18
        NightThreshold      = 23
)

// message constants
const (
        DefaultMessage      = "! Welkom bij Fonteyn Vakantieparken"
        MorningPrefix       = "Goedemorgen"
        AfternoonPrefix     = "Goedemiddag"
        EveningPrefix       = "Goedenavond"
        MorningMessage      = MorningPrefix + DefaultMessage
        AfternoonMessage    = AfternoonPrefix + DefaultMessage
        EveningMessage      = EveningPrefix + DefaultMessage
        NightMessage        = "Sorry, de parkeerplaats is ’s nachts gesloten"
        InteractivePrompt   = "Kentekenplaat (of '>EXIT' om af te sluiten): "
        AccessDeniedMessage = "U heeft helaas geen toegang tot het parkeerterrein"
)

// license plates
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

//flag variables
var (
    debug bool
    debugTime time.Time
    interactive bool
)

func init() {
    // parse flags
    flag.BoolVar(&debug,"debug", false, "Toggles debug mode")
    flag.BoolVar(&debug, "d", false, "Toggles debug mode (shorthand)")

    flag.TextVar(&debugTime,"debugOffset", time.Time{}, "debug offset")
    flag.TextVar(&debugTime,"o", time.Time{}, "debug offset (shorthand)")

    flag.BoolVar(&interactive,"interactive", false, "Toggles interactive mode")
    flag.BoolVar(&interactive, "i", false, "Toggles interactive mode (shorthand)")

    flag.Parse()
}



func main() {
    // prepare exit code
    exitCode := 0
    defer func() {os.Exit(exitCode)}()

    // parse variables based on flags
    var (
        
        now = func() time.Time {
            if debug {
                return debugTime
            } else {
                return time.Now()
            }
        }()

        reader = func() *bufio.Reader{
            if interactive {
                return bufio.NewReader(os.Stdin)
            } else {
                return bufio.NewReader(strings.NewReader(strings.Join(flag.Args(), "\n")))
            }
        }()

        licensePlates = func() TLicensePlates {
            if debug {
                return Set.NewSet(debugLicensePlates()...)
            } else {
                // PROD: get license plates from database
                return Set.NewSet(debugLicensePlates()...)
            }
        }()
    )

    // read license plates
    var (licensePlate string; err error)
    for err == nil {
        if interactive {fmt.Print(InteractivePrompt)}
        var licensePlateBytes []byte;
        licensePlateBytes ,_, err = reader.ReadLine()
        licensePlate = string(licensePlateBytes)
        
        // check if interactive mode exit is requested
        if interactive&&(licensePlate == ">EXIT") || err != nil {
            break
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
            fmt.Println(AccessDeniedMessage)
        }
    }
    if err != nil && err != io.EOF {
        exitCode = 1
    }
}
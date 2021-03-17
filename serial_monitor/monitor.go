// Package main  This is not a tinygo project!
// This is meant to be used as the serial monitor (like the one arduino has!)
// You should be able to read print() and println() calls while running this program
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tarm/serial"
)

// edit this according to your Arduino's port and the file you desire to print to
var config = struct {
	Port     string
	Filename string
}{
	// the port will be COMx on windows (where x is a number).
	// on linux the port will take the form /dev/ttyACM0 or /dev/ttyUSB0 (or similar)
	Port:     "/dev/ttyACM0",
	Filename: "outduino.txt",
}

func main() {
	if len(os.Args) > 1 {
		config.Port = os.Args[1]
	}
	c := &serial.Config{Name: config.Port, Baud: 9600}
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}
	fo, err := os.Create(config.Filename)
	defer fo.Close()
	if err != nil {
		panic(err)
	}

	var n int
	buf := make([]byte, 128)
	for {
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
		fmt.Printf("%s", buf[:n])
	}
}

// Package main  This is not a tinygo project!
// This is meant to be used as the serial monitor (like the one arduino has!)
// You should be able to read print() and println() calls while running this program
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/tarm/serial"
)

// edit this according to your Arduino's port and the file you desire to print to
var config = struct {
	Port     string
	Filename string
	Baud     int
}{
	// the port will be COMx on windows (where x is a number).
	// on linux the port will take the form /dev/ttyACM0 or /dev/ttyUSB0 (or similar)
	Port:     "/dev/ttyACM0",
	Filename: "",
	Baud:     9600,
}

func main() {
	pflag.IntVarP(&config.Baud, "baud", "b", config.Baud, "Baudrate for port. Common values: 9600, 115200")
	pflag.StringVarP(&config.Filename, "output", "o", config.Filename, "File to pipe read data to. Data is still piped to stdout.")
	pflag.Parse()

	if len(pflag.Args()) >= 1 {
		config.Port = pflag.Arg(0)
	}
	c := &serial.Config{Name: config.Port, Baud: config.Baud}
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}
	fo := io.Discard
	if config.Filename != "" {
		fo, err := os.Create(config.Filename)
		defer fo.Close()
		if err != nil {
			panic(err)
		}
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

// Package lia_lcd shows simple use of tinygo
// to read Arduino Uno A0 pin and print value to LCD screen (hd44780, 4 bit config).
// The arduino LED will blink as well.
//
// LCD connection scheme is same as shown at https://www.arduino.cc/en/Tutorial/LibraryExamples/HelloWorld
//
// E: Pin11, RS: Pin12,  RW: NoPin
package main

import (
	"machine"
	"strconv"
	"time"

	// If tinygo errors when flashing due to package import error, install Go version 1.15.8, run go mod list (generate the go mod file) and reinstall Go 1.16 to be able to flash
	"tinygo.org/x/drivers/hd44780"
)

const maxUint16 uint16 = 0xffff

// Arduino pins to LCD Dx pins 4 thru 7
var Dpins = [4]machine.Pin{
	// Pins 4..7 on LCD are hooked to pins 5..2 on arduino board.
	machine.D5, machine.D4, machine.D3, machine.D2,
}

// Potentiometer pin shall be configured to be an ADC. We use pin A0
var pote = machine.ADC{
	Pin: machine.ADC0,
}

func main() {
	// ADC configuration
	machine.InitADC()
	adcCfg := machine.ADCConfig{}
	pote.Configure(adcCfg)

	// LED will blink. Thus we configure it to be an output.
	machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// LCD configs
	lcd, _ := hd44780.NewGPIO4Bit(
		Dpins[:],
		machine.D11, machine.D12, machine.NoPin,
	)
	lcd.Configure(hd44780.Config{
		Width:       16,
		Height:      2,
		CursorOnOff: false,
		CursorBlink: false,
	})

	// Write tool versions at the time of writing this program
	lcd.ClearDisplay()
	lcd.Write([]byte("TinyGo0.17 Go16"))
	lcd.Display()

	var val uint8
	for {
		// set line place to overwrite previous content
		lcd.SetCursor(0, 1)
		// read and print adc value
		val = adcToPcnt(pote.Get())
		lcd.Write([]byte(strconv.Itoa(int(val))))
		lcd.Display()
		// Blinky. On Arduino uno machine.LED is pin 13. (machine package specifies microchip pins, these may or may not coincide with board pin names)
		machine.LED.High()
		time.Sleep(time.Millisecond * 200)
		machine.LED.Low()
		time.Sleep(time.Millisecond * 200)
	}
}

// Converts ADC uint16 value to a percent (uint8).
// As of creating this program there is no full support of
// 32bit floating point operations on AVR boards (Atmega boards mainly)
func adcToPcnt(v uint16) uint8 {
	const div uint16 = maxUint16 / 100
	return uint8(v / div)
}

// Package lia_lcd shows simple use of tinygo
// to read Arduino Uno A0 pin and print value to LCD screen (hd44780, 4 bit config) and create
// a PWM signal on pin D6.
// We also read A1 and print value read to lcd.
// The arduino LED will blink as well.
//
// LCD connection scheme is same as shown at https://www.arduino.cc/en/Tutorial/LibraryExamples/HelloWorld
//
// E: Pin11, RS: Pin12,  RW: NoPin
package main

import (
	"machine"
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

var sensor = machine.ADC{
	Pin: machine.ADC1,
}

const (
	pwmPin       = machine.D6
	blinkyPeriod = time.Millisecond * 500
)

func main() {
	// PWM Config
	machine.InitPWM()
	pwm := machine.PWM{Pin: pwmPin}
	err := pwm.Configure()
	checkError(err, "failed to configure pwmPin")
	// ADC configuration
	machine.InitADC()
	adcCfg := machine.ADCConfig{}
	pote.Configure(adcCfg)
	sensor.Configure(adcCfg)
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
	// lcd.Write([]byte("    bar"))
	lcd.SetCursor(0, 0)
	lcd.Write([]byte("ADC"))
	lcd.Display()
	lcd.SetCursor(0, 1)
	lcd.Write([]byte("PWM"))

	lcd.Display()
	// go blinky()
	var val uint8
	// ton, toff := millis(), millis()
	var numbuff = [3]byte{'0', '0', '0'}
	for {
		// PWM control through potentiometer
		// set line place to overwrite previous content
		lcd.SetCursor(5, 1)
		// read and print adc value
		val = adcToPcnt(pote.Get())
		formatUint8(val, numbuff[:])
		lcd.Write(numbuff[:])
		pwmval := (uint16(val) * 254) / 100
		pwm.Set(pwmval << 8)
		lcd.Display()
		// analogRead on pin A1
		val = adcToPcnt(sensor.Get())
		lcd.SetCursor(5, 0)
		formatUint8(val, numbuff[:])
		lcd.Write(numbuff[:])
		lcd.Display()
	}
}

// Converts ADC uint16 value to a percent (uint8).
// As of creating this program there is no full support of
// 32bit floating point operations on AVR boards (Atmega boards mainly)
func adcToPcnt(v uint16) uint8 {
	const div uint16 = maxUint16 / 100
	return uint8(v / div)
}

func checkError(err error, msg string) {
	if err != nil {
		print(msg, ": ", err.Error())
		println()
	}
}

var tstart = time.Now()

func millis() int {
	return int(time.Since(tstart).Milliseconds())
}

var str string

func formatUint8(val uint8, buff []byte) {
	buff[2] = val%10 + '0'
	buff[1] = (val/10)%10 + '0'
	buff[0] = (val/100)%10 + '0'
}

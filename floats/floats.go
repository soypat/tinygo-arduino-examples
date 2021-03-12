package main

// On the Arduino Uno the FPU works with 32 bit floats
// This means that the "math" package will not work on the arduino Uno!
// To use math functions defer to github.com/chewxy/math32 package.
// Do note that most functions in math32 are just wrappers around
// 64 bit stdlib functions and will not work.
type float float32

func main() {
	var v float
	v = 2. / 3.
	println(v) // prints to Serial port.
}

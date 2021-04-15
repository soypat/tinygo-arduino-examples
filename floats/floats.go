// Package main On the Arduino Uno the FPU works with 32 bit floats
// This means that the "math" package will not work on the arduino Uno!
// To use math functions defer to github.com/chewxy/math32 package.
// Do note that most functions in math32 are just wrappers around
// 64 bit stdlib functions and will not work.
package main

// Type alias for working with floats. Nice shorthand to float32
type float = float32

// Note creating a user defined type will mean you will not be able to
// pass `userfloat` types to functions that take a float32. Avoid this pattern if possible
type userfloat float32

func main() {
	var v float
	v = 2. / 3.
	println(v) // prints to Serial port.

	// Go's default float type is float64. One must always
	// have the type be explicit so the code compiles
	pi := float(3.14159)
	println(pi)

	// A one liner to declare and initialize a float
	var c float = 22.123
	// the compiler knows 23.12 is a float32
	pi = c * 23.12
	println(c)

	var u uint16 = 256
	println(float(u)) // cast integer to float

}

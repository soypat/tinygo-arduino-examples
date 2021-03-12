# Floating point operations

Tinygo uses the LLVM toolchain to create programs. For now, only 32 bit floats are supported on AVR boards by LLVM. This comes into conflict with Go's default 64 bit floats. 

To be able to compile to arduino boards the `float32` type must be explicit! Trying to initialize a float implicitly will yield a compile-time error. i.e:

```go
// WILL NOT COMPILE
myfloat := 3.1415 // typed as float64 implicitly
println(myfloat)
```

The following program will compile correctly:
```go
var myfloat float32 = 3.1415
println(myfloat)
```


To get around this elegantly one can create a [type alias](https://yourbasic.org/golang/type-alias/)

```go
type float = float32

var v float = 3.1415
myfloat := float(22.13)
println(v*myfloat*2.712)
```

using type aliases one can switch to `float64` easily for when they are eventually supported.

## Casting

To cast an integer to a float do it as you would in Go

```go
var adcValue uint16 = 2030

println(float32(adcValue*20))
```

If you have defined a type alias, such as `type float = float32` then one can cast using `float(adcValue)`


## Converting to string
You'll have to write your own function. Arduino Uno's memory is very limited and the standard library `strconv.FormatFloat` is enormous (to be fair, [it does A LOT](https://golang.org/pkg/strconv/#FormatFloat)), not to mention it requires a `float64`.

Your best bet is using a [`modf`](https://golang.org/src/math/modf.go) styled function and then converting the integer and fraction part of the float to a string using `stconv.Itoa`.

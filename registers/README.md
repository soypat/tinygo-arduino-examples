# Registers

There may be slight differences in how register flags are included in arduino vs tinygo. 
For example, to set the `SPCR` registry (arduino SPI clock) one in arduino can simply write

```cpp
void setup() {
    // arduino setting of SPCR registry 0x4c (77)
    SPCR = bit(SPE) | bit(MSTR);
    bitSet(SPSR, SPI2X);
}
```
SPE and MSTR are bitflags 6 and 4, respectively, so they must be set with the `bit()` function.

In tinygo we import the `device/avr` package to have access to registry addresses, bitmasks and flags along with
several registry modifying primitives. To do the same thing in tinygo we write:

```go
package main

import "device/avr"

func main() {
    avr.SPCR.Set(avr.SPCR_SPE | avr.SPCR_MSTR)
    avr.SPSR.SetBits(avr.SPSR_SPI2X)
}
```
 `avr.SPCR_SPE` and `avr.SPCR_MSTR` are the equivalent of arduino's `bit(SPE)` and `bit(MSTR)` ,respectively. It is preffered to avoid bitflags in tinygo and always use values that only require basic bitwise operations like `&`,  `|` and `^` (NOT, equivalent to C's `~`).
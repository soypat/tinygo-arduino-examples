# Optimization
you won't need to...
---

We may sometimes be constrained by memory or time limitation, for this compile flags are your friend. 

The following two examples yield the same bit-code

```go
func setBank(address uint8) {
    bank := address & BANK_MASK // we create a variable in function scope
	if bank != d.Bank {
		d.writeOp(BIT_FIELD_CLR, ECON1, bank)
		d.writeOp(BIT_FIELD_SET, ECON1, bank>>5)
		d.Bank = bank
	}
}
```
One could argue such a variable `bank` is not needed and that we lose precious memory space. 

This could not be farther from the truth. LLVM already has an aggressive optimizer implemented which TinyGo 
uses by default (`-opt=z` flag). The following code may look like it would consume less space, however LLVM produces
the same bitcode for both functions.

```go
func setBank(address uint8) {
	if address & BANK_MASK != d.Bank {
		d.writeOp(BIT_FIELD_CLR, ECON1, address & BANK_MASK)
		d.writeOp(BIT_FIELD_SET, ECON1, (address & BANK_MASK)>>5)
		d.Bank = address & BANK_MASK
	}
}
```

Hopefully this will clear up some misconceptions on optimizing for space: **it is preferrable
to write readable code** from the start and let the compiler optimize as much as it can. Once
you are done with your program and you find it runs too slow or could consume less space, then one begins optimizing where convenient.
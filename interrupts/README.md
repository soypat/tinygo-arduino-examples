# Interrupts

The arduino library has two widely known functions `noInterrupts()` and `interrupts()` to control the firing of interrupts in critical sections of code. Beneath the hood these are assembly instructions like so:

```cpp
// this is noInterrupts()
# define cli()  __asm__ __volatile__ ("cli" ::: "memory")

// this is interrupts()
# define sei()  __asm__ __volatile__ ("sei" ::: "memory")
```
both `cli()` and `sei()` can be used in arduino code.

In TinyGo we have finer control of interrupts using the `runtime/interrupt` library.

```go
import "runtime/interrupt"

func main() {
    state := interrupt.Disable()
    // what you want to do with interrupts disabled
    interrupt.Restore(state)
    // interrupts are enabled again (if they were enabled before `interrupt.Disable()`
}
```

This allows us to nest critical code without fear of accidentally clearing an interrupt in the calling function. `state := interrupt.Disable()` is equal to `noInterrupts()` and `interrupt.Restore(state)` is equal to `interrupts()`. This prevents hard to debug problems that would happen if we were using `cli()` and `sei()`.

Below is an example of a program that is unsafe to use with `cli()` and `sei()` but is safe to use with the `interrupt` library.

```go
func foo() {
    state := interrupt.Disable()
    // do something with interrupts disabled
    interrupt.Restore(state)
}
func main() {
    state := interrupt.Disable()
    // do something with interrupts disabled
    foo()
    // do something more, with interrupts still disabled (!)
    interrupt.Restore(state)
}
```


`State` should be treated as an opaque number. On AVR `state` contains the `SREG` register while on other platforms it could be `0` or `1` depending on whether interrupts were previously disabled or not. 




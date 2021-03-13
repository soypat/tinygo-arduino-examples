# Size of TinyGo Programs

We may be familiar with the large binaries of Go programs. Your typical Go helloworld occupies 1.2MB of space on your disk drive. That's a whole lot for just printing `Hello World`! Certainly too large for microcontrollers. This is because Go embeds the whole runtime library (goroutines, garbage collector and scheduler).

TinyGo remedies this by using LLVM toolchain to compile. You can build the same program for your computer (if running linux) with the `tinygo` command:

```console
tinygo build -o="tinyhello" .
```

Running `ls -lah` shows us the size of the programs we compiled:

```console
> ls -lah
total 1,3M
drwxrwxr-x 2 pato pato 4,0K mar 13 11:29 .
drwxrwxr-x 9 pato pato 4,0K mar 13 11:19 ..
-rw-rw-r-- 1 pato pato   22 mar 13 11:19 go.mod
-rwxrwxr-x 1 pato pato 1,2M mar 13 11:20 hello
-rw-rw-r-- 1 pato pato   54 mar 13 11:19 helloworld.go
-rw-rw-r-- 1 pato pato  451 mar 13 11:28 README.md
-rwxrwxr-x 1 pato pato  22K mar 13 11:29 tinyhello
```

The size of the linux binary is only 22 kilobytes! Down from 1200!

This size is still too large for microcontrollers though, especially the Arduino UNO which lives on 32kB of flash memory. To build the binary that is flashed to the arduino one can specify the `-target` and build the same file

```console
tinygo build -target=arduino -o="tinyhello.bin" .
```

The `.bin` extension [is important for the tinygo `build` command](https://tinygo.org/usage/subcommands/), it specifies the output file type. Running `ls -lah` again shows us the `hello.bin` file is **625 bytes**, about the size one would expect for such a simple program.

To view more information on program size use the `-size=short` flag with `tinygo build`

## Optimization


One can further reduce the size of the binary by discarding features one does not use. See important [build options on the tinygo.org site.](https://tinygo.org/usage/important-options/).

### RAM vs. Flash
By using the `-size=short` flag to view memory usage we find if we use a less aggressive optimization strategy (`-opt=s` instead of `-opt=z`) we reduce our ram usage by 81 bytes while increasing flash storage usage by 14 bytes. Usually flash memory is more critical, but it is nice to be able to manage both.

### Garbage collector

Another factor one has control over is the **garbage collector**. One can remove the garbage collector entirely from the program using the flag `-gc=none`. The program in this case is reduced by **146 bytes** hitting 479 kB of flash memory. *This flag will cause a compile error if memory is allocated*.

Usually all TinyGo programs run with the garbage collector enabled and the reason the `none` options is available is to find places where the program allocates memory.

### Panic

Panic handling consumes flash memory so as to print expressive error messages. Sometimes these messages will not be viewable, so it makes sense to disable this. `-panic=trap` reduced flash usage by **82 bytes**.

### Scheduler
If goroutines and channels are not needed, one can completely remove the scheduler on TinyGo programs with the flag `-scheduler=none`.

This **did not reduce the size of the program** as it is enabled by default. If `-schedulers=coroutines` then binary size increases by 32 bytes and RAM increases as well.

### Sum of all parts
```console
$ tinygo build -target=arduino -size=short -o="hello.bin"  .
   code    data     bss |   flash     ram
    544      81     649 |     625     730

$ tinygo build -target=arduino -size=short -o="hello.bin" -gc=none -panic=trap  .
   code    data     bss |   flash     ram
    456      11     642 |     467     653
```
Enabling `opt=z`, panic and garbage collector optimizations a **467 kB** binary is reached, down from **625 kB**.

## Conclusion
Running `tinygo build` can show us the size the program will have on the microcontroller. We also saw that we can reach much smaller sizes if we use tinygo instead of the go compiler for regular computer programs!
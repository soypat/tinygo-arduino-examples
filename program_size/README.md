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

## Conclusion
Running `tinygo build` can show us the size the program will have on the microcontroller so that we don't bloat it. We also saw that we can reach much smaller sizes if we use tinygo instead of the go compiler for regular computer programs!
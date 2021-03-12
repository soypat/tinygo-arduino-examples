# tinygo-arduino-examples
Compilation of LCD screen, ADC, and output examples.

## It is suggested you run tinygo with Go 1.15 until [#44557](https://github.com/golang/go/issues/44557) is fixed

## Run an example
To run any example navigate to the example directory in command line

```console
cd lcdscreen_adc
```
Then flash the arduino. You must know the port it's on. On windows they typically are of the form `COM1` through `COM4`. On Linux the port may look something like `/dev/ttyACM0` or `/dev/ttyUSB0`.

```console
tinygo flash -target=arduino -port=/dev/ttyUSB0 .
```

If everything goes OK, some progress bars should pop up and the process should end with a thankful message:

```console
avrdude: verifying ...
avrdude: 7036 bytes of flash verified

avrdude done.  Thank you.
```

You now have tinygo running on your Arduino!

## To create your own "sketch"

1. Create a directory and create a `.go` file of any name. Say we make `twister.go`. 
2. Create the module. Run `go mod init my_awesome_module` (with any module name you find appropiate for your program)
3. Start coding in this directory. You may create more files if you wish to distribute you functions
    * To enable intellisense on VSCode follow the [instructions on IDE integration on tinygo.org](https://tinygo.org/ide-integration/). You basically have to create the same [`.vscode/settings.json`](.vscode/settings.json) file this repo has but with you own `GOROOT` Path (this may differ between installations)
4. Run `go mod tidy` when done programming. And flash your sketch with 

```console
tinygo flash -target=arduino -port=/dev/ttyUSB0 .
```
# Serial monitor

This is ***not*** a tinygo project!

This program is meant to be used as the serial monitor (like the one arduino has!)

You should be able to read print() and println() calls from the arduino while running this program on your computer (connected to arduino via USB)

### Instructions
To run this program, write in console:

```console
go run .
```
This will wait until it recieves a message on the port and then shall print the message. This will repeat until the console is closed or you press <kbd>Ctrl</kbd>+<kbd>C</kbd>.

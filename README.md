# digispark-go

This project is inspired by the Digispark's sample code:
https://github.com/digistump/DigisparkExamplePrograms/blob/master/Python/DigiUSB/source/arduino/usbdevice.py

## Usage

First, install libusb:

```sh
# Ubuntu22.04
sudo apt install libusb-1.0-0-dev
```

Build and send data to digispark board:

```sh
make

sudo ./bin/digispark write 1  # turn on LED
sudo ./bin/digispark write b  # blink LED
sudo ./bin/digispark write 0  # turn off LED
```

## Sample Digispark Code

Write firmware to the Digispark board:

```sh
code misc/digispark
# THis is a PlatformIO-based project.
# You need the PlatformIO extension or vanila PlatformIO
# https://docs.platformio.org/en/latest/integration/ide/vscode.html#quick-start
```

1. Unplug a digispark board if connected
2. Press Upload button
3. Plug in the digispark board when the message `Please plug in the device` is shown

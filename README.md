# digispark-go

This project is inspired by the Digispark sample code:
https://github.com/digistump/DigisparkExamplePrograms/blob/master/Python/DigiUSB/source/arduino/usbdevice.py

## Overview

digispark-go is a simple tool to communicate with a Digispark board from your Linux PC using libusb. You can control the onboard LED or send custom data to your own Digispark firmware.

## Prerequisites

- Linux (tested on Ubuntu 22.04)
- Digispark board (ATtiny85)
- [libusb-1.0-0-dev](https://libusb.info/)

Install libusb:

```sh
# Ubuntu22.04
sudo apt install libusb-1.0-0-dev
```

## Build & Usage

Build the tool and send commands to your Digispark board:

```sh
make

# Turn on LED
sudo ./bin/digispark write 1
# Blink LED
sudo ./bin/digispark write b
# Turn off LED
sudo ./bin/digispark write 0
```

## Writing Firmware for Digispark

The `misc/digispark` directory contains a PlatformIO-based firmware project for the Digispark board.

To open and upload the firmware:

```sh
code misc/digispark
# This is a PlatformIO-based project.
# You need the PlatformIO extension or vanilla PlatformIO.
# https://docs.platformio.org/en/latest/integration/ide/vscode.html#quick-start
```

1. Unplug the Digispark board if connected.
2. Press the Upload button in VSCode.
3. Plug in the Digispark board when prompted with `Please plug in the device`.

---

Feel free to customize the firmware in `misc/digispark/src/main.cpp` to suit your needs. For more information, see the original Digispark documentation and sample code linked above.

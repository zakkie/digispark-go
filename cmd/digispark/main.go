package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zakkie/digispark-go/internal/digiusb"
)

type ErrInvalidCommand struct {
	Message string
}

func (e *ErrInvalidCommand) Error() string {
	return fmt.Sprintf("invalid command: %s", e.Message)
}

func usage() {
	fmt.Println("Usage: digispark <read|write> [data]")
	fmt.Println("  read: Read data from the device")
	fmt.Println("  write <data>: Write data to the device")
}

func main() {
	err := _main()
	if err != nil {
		log.Printf("Error: %s\n", err)
		switch err.(type) {
		case *ErrInvalidCommand:
			usage()
		}
		os.Exit(1)
	}
}

func _main() error {
	if len(os.Args) < 2 {
		return &ErrInvalidCommand{Message: "no command provided"}
	}

	usb, err := digiusb.New()
	if err != nil {
		return fmt.Errorf("error opening device: %w", err)
	}
	defer usb.Close()

	command := os.Args[1]
	switch command {
	case "read":
		data, err := usb.Read()
		if err != nil {
			return &ErrInvalidCommand{Message: "error reading from device"}
		}
		log.Printf("Read data: %s\n", string(data))

	case "write":
		if len(os.Args) < 3 {
			return &ErrInvalidCommand{Message: "no data provided for write command"}
		}
		data := os.Args[2]

		for _, b := range data {
			err := usb.Write(byte(b))
			if err != nil {
				return fmt.Errorf("error writing to device: %w", err)
			}
		}
		log.Printf("Wrote data: %s\n", data)

	default:
		return &ErrInvalidCommand{Message: fmt.Sprintf("unknown command: %s", command)}
	}

	return nil
}

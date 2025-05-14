package digiusb

import (
	"fmt"

	"github.com/google/gousb"
	"github.com/google/gousb/usbid"
)

type DigiUSB struct {
	ctx *gousb.Context
	dev *gousb.Device
}

const (
	vendorId  gousb.ID = 0x16c0 // Digispark vendor ID
	productId gousb.ID = 0x05df // Digispark product ID
)

const (
	ctrlOut             uint8 = 0x00
	ctrlIn              uint8 = 0x80
	ctrlTypeClass       uint8 = 1 << 5
	ctrlRecipientDevice uint8 = 0x00
)

const (
	usbReqHidSetReport      uint8  = 0x09
	usbHidReportTypeFeature uint16 = 0x03
)

func New() (*DigiUSB, error) {
	ctx := gousb.NewContext()
	dev, err := ctx.OpenDeviceWithVIDPID(vendorId, productId)
	if err != nil || dev == nil {
		ctx.Close()
		return nil, fmt.Errorf("device not found: %w", err)
	}

	// enable auto detach explicitly
	dev.SetAutoDetach(true)
	return &DigiUSB{ctx: ctx, dev: dev}, nil
}

func (d *DigiUSB) Close() {
	d.dev.Close()
	d.ctx.Close()
}

// Write a byte to the device
func (d *DigiUSB) Write(data byte) error {
	reqType := ctrlRecipientDevice | ctrlTypeClass | ctrlOut
	_, err := d.control(reqType, usbReqHidSetReport, (usbHidReportTypeFeature << 8), uint16(data), nil)
	return err
}

// Read a byte from the device
func (d *DigiUSB) Read() (byte, error) {
	panic("TODO: not implemented")
}

func (d *DigiUSB) String() string {
	return fmt.Sprintf("DigiUSB: %s - %s", d.dev, usbid.Describe(d.dev.Desc))
}

// Internal functions
func (d *DigiUSB) control(requestType, request uint8, value, index uint16, data []byte) (int, error) {
	return d.dev.Control(requestType, request, value, index, data)
}

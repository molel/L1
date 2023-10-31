package main

import "fmt"

type UsbTypeA struct {
}

func (u *UsbTypeA) InsertUsbTypeAIntoComputer(computer Computer) {
	computer.InsertUsbTypeA()
}

type Computer interface {
	InsertUsbTypeA()
}

type ComputerWithUsbTypeA struct {
}

func (c *ComputerWithUsbTypeA) InsertUsbTypeA() {
	fmt.Println("Inserting USB Type A into computer")
}

type ComputerWithUsbTypeC struct {
}

func (c *ComputerWithUsbTypeC) InsertUsbTypeC() {
	fmt.Println("Inserting USB Type C into computer")
}

type UsbTypeAAdapter struct {
	c *ComputerWithUsbTypeC
}

func (a *UsbTypeAAdapter) InsertUsbTypeA() {
	fmt.Println("Inserting adapter (USB Type C -> USB Type A)")
	a.c.InsertUsbTypeC()
}

func main() {
	// client
	usbTypeA := UsbTypeA{}

	// concrete prototype
	computerWithUsbTypeA := ComputerWithUsbTypeA{}

	// adaptee
	computerWithUsbTypeC := ComputerWithUsbTypeC{}
	usbTypeAAdapter := UsbTypeAAdapter{c: &computerWithUsbTypeC}

	usbTypeA.InsertUsbTypeAIntoComputer(&computerWithUsbTypeA)
	usbTypeA.InsertUsbTypeAIntoComputer(&usbTypeAAdapter)

}

//Command interface
package main

type Command interface {
	execute()
}
type OnCommand struct {
	device Device
}
type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

func (c *OnCommand) execute() {
	c.device.on()
}

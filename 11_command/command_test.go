package command

func ExampleCommand() {
	mb := &MotherBoard{}
	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)

	box1 := NewBox(startCommand, rebootCommand)
	box1.PressButton1()
	box1.PressButton2()

	box2 := NewBox(rebootCommand, startCommand)
	box2.PressButton1()
	box2.PressButton2()
	// Output:
	// system starting
	// system rebooting
	// system rebooting
	// system starting
}

package gadb

func (d Device) Home() {
	d.RunShellCommand("input", "keyevent", "3")
}

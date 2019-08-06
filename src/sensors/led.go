package sensors

type led struct {
	pinNumber pi.Pin
	isOn      bool
}

func NewLED(p pi.Pin) *led {
	return &led{
		pinNumber: p,
		isOn:      false,
	}
}

func (l *led) Toggle() {
	l.pinNumber.Toggle()
	l.isOn = !l.isOn
}

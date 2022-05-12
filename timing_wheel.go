package timingwheel

type TimingWheel struct {
}

func NewTimingWheel() *TimingWheel {
	return &TimingWheel{}
}

func (tw *TimingWheel) Start() {
}

type Task interface {
}

func (tw *TimingWheel) Add(t Task) error {
	return nil
}

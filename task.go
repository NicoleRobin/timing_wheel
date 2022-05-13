package timingwheel

type Task interface {
	Execute(t Task) error
}

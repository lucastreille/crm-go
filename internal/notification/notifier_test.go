package notification

import "testing"

type FakeNotifier struct {
	Calls int
}

func (f *FakeNotifier) Send(message string) error {
	f.Calls++
	return nil
}

func TestNotifyAll(t *testing.T) {
	n1 := &FakeNotifier{}
	n2 := &FakeNotifier{}

	notifiers := []Notifier{n1, n2}

	NotifyAll(notifiers, "hello world")

	if n1.Calls != 1 {
		t.Errorf("n1.Calls = %d, want 1", n1.Calls)
	}

	if n2.Calls != 1 {
		t.Errorf("n2.Calls = %d, want 1", n2.Calls)
	}
}

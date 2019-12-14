package task

/**
 * Interface wanted for the Network
 */
type Network interface {
	EmitEcho(id uint16) bool
}

type Manager interface {
	GetElected() uint16
}

/**
 * Represents an applicative Task
 */
type Task struct {
	currentElected uint16
	shouldRunElection bool
	m Manager
	n Network
}

func (t *Task) Run() {
	t.shouldRunElection = true

	for {
		if t.shouldRunElection {
			t.currentElected = t.m.GetElected()
			t.shouldRunElection = false
		} else { // TODO is it correct?
			if t.n.EmitEcho(t.currentElected) == false {
				t.shouldRunElection = true
			}
		}
	}
}
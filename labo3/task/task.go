package task

import (
	"log"
	"time"
)

/**
 * Interface wanted for the Network
 */
type Network interface {
	EmitEcho(id uint16) bool
}

type Manager interface {
	GetElected() uint16
	RunElection()
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

func (t *Task) Run(manager Manager, network Network) {
	log.Println("Task : Initialization of the task")
	t.m = manager
	t.n = network
	t.shouldRunElection = true

	t.m.RunElection()

	for {

		log.Println("Task : get the elected processus")
		t.currentElected = t.m.GetElected()
		hasAnswered := t.n.EmitEcho(t.currentElected)
		if !hasAnswered {
			t.m.RunElection()
		}
		time.Sleep(time.Second * 1)
	}
/*
	for {
		if t.shouldRunElection {
			log.Println("Task : get the elected processus")
			t.currentElected = t.m.GetElected()
			t.shouldRunElection = false
		} else { // TODO is it correct?
		time.Sleep(time.Second * 1)
			log.Println("Task : Emit Echo")
			if t.n.EmitEcho(t.currentElected) == false {
				t.shouldRunElection = true
			}
		}
	}
*/
}
/*
 -----------------------------------------------------------------------------------
 Lab 		 : 03
 File    	 : config.go
 Authors   	 : François Burgener - Tiago P. Quinteiro
 Date        : 10.12.19

 Goal        : TODO
 -----------------------------------------------------------------------------------
*/
package manager

import (
	"log"
)

/**
 * ENUM declaration of the states
 */
const (
	REST = iota
	NOTIFICATION
	RESULT
)

/**
 * Interface wanted for the Network
 */
type Network interface {
	EmitNotif(map[uint16]uint16)
	EmitResult(uint16, map[uint16]bool)
}

type ResultMessage struct {
	id uint16
	visitedResult map[uint16]bool
}

type Manager struct {
	N uint16
	me uint16
	aptitude uint16
	state uint8 // TODO Maybe change this
	elected uint16
	network Network
	chanAskElection chan bool
	chanGiveElection chan uint16
	chanNotification chan map[uint16]uint16
	chanResult chan ResultMessage
}

func (m *Manager) Init(N uint16, me uint16, aptitude uint16, network Network) {
	log.Println("Manager : Initialization of the manager")
	m.N = N
	m.me = me
	m.aptitude = aptitude
	m.network = network
	m.state = REST

	//Channels
	m.chanAskElection = make(chan bool)
	m.chanGiveElection = make(chan uint16)
	m.chanNotification = make(chan map[uint16]uint16)
	m.chanResult = make(chan ResultMessage)


	go m.handler()
}


func (m *Manager) handler() {
	for {
		select {
		case <- m.chanAskElection:
			l := make(map[uint16]uint16)
			l[m.me] = m.aptitude

			log.Println("Manager : Emit notification")
			m.network.EmitNotif(l)
			m.state = NOTIFICATION
		case notifMap := <- m.chanNotification:
			log.Println("Manager : Received notification : ",notifMap, " me:",m.me)
			_, isInside := notifMap[m.me] // Test if I'm here
			if isInside {
				m.elected = findMax(notifMap)

				resultMap := make(map[uint16]bool)
				resultMap[m.me] = true // TODO Could be void struct.

				m.network.EmitResult(m.elected,resultMap)
				m.state = RESULT
			} else {
				notifMap[m.me] = m.aptitude // Add myself in map
				m.network.EmitNotif(notifMap)
				m.state = NOTIFICATION
			}
		case resultMessage := <- m.chanResult:
			i := resultMessage.id
			resultMap := resultMessage.visitedResult

			_, isInside := resultMap[m.me] // Test if I'm here
			if isInside {
				// Nothing to do ¯\_(ツ)_/¯
			} else if m.state == RESULT && m.elected != i {
				// TODO this code is similar to another
				l := make(map[uint16]uint16)
				l[m.me] = m.aptitude

				m.network.EmitNotif(l)
				m.state = NOTIFICATION
			} else if m.state == NOTIFICATION {
				m.elected = i

				// TODO this code is similar to another
				resultMap := make(map[uint16]bool)
				resultMap[m.me] = true // TODO Could be void struct.

				m.network.EmitResult(m.elected,resultMap)
				m.state = RESULT
			}
		default:
			if m.state == RESULT {
				log.Println("Manager : Send elected processus")
				m.chanGiveElection <- m.elected
			}
		}
	}
}

// API

/**
 * Submits a Notification message to manager from network
 */
func (m *Manager) SubmitNotification(notifMap map[uint16]uint16) {
	m.chanNotification <- notifMap
}

/**
 * Submits a result message to manager from network
 */
func (m *Manager) SubmitResult(id uint16, resultMap map[uint16]bool) {
	m.chanResult <- ResultMessage{
		id:            id,
		visitedResult: resultMap,
	}
}

/**
 * Get the elected id
 */
func (m *Manager) GetElected() uint16 {
	log.Println("Manager : get the elected processus")
	m.startElection()
	return <- m.chanGiveElection
}

/**
 * Tells manager to start an election
 */
func (m *Manager) startElection(){
	log.Println("Manager : Start election")
	m.chanAskElection <- true
}


/**
 * @param m Map where you want to find max
 */
func findMax (m map[uint16]uint16) uint16 {
	var id, max uint16 = 0, 0

	for key, val := range m {
		if val > max {
			max = val
			id = key
		}
	}

	return id
}
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

/**
 * ENUM declaration of the states
 */
const (
	REST = iota
	BUSY
	NOTIFICATION
	RESULT
)

type Void struct {}

/**
 * Interface wanted for the Network
 */
type Network interface {
	EmitNotif(map[uint16]uint16)
	EmitResult(uint16, map[uint16]Void)
}

type Manager struct {
	N uint16
	me uint16
	aptitude uint16
	state uint8 // TODO Maybe change this
	elected uint16
	network Network
	chanElection chan bool
	chanNotification chan map[uint16]uint16
	chanResult chan map[uint16]Void
}

func (m *Manager) Init() {
	m.state = REST
	go m.handler()
}


func (m *Manager) handler() {
	for {
		select {
		case <- m.chanElection:
			l := make(map[uint16]uint16)
			l[m.me] = m.aptitude

			m.network.EmitNotif(l)
			m.state = NOTIFICATION
		case notifMap := <- m.chanNotification:
			_, isInside := notifMap[m.me] // Test if I'm here
			if isInside {
				m.elected = findMax(notifMap)

				resultMap := make(map[uint16]Void)
				resultMap[m.me] = Void{}

				m.network.EmitResult(m.elected,resultMap)
				m.state = RESULT
			} else {
				notifMap[m.me] = m.aptitude // Had myself in map
				m.network.EmitNotif(notifMap)
				m.state = NOTIFICATION
			}
		case resultMap := <- m.chanResult:
			// TODO
			_, isInside := resultMap[m.me] // Test if I'm here
			if isInside {

			} else if m.state == NOTIFICATION{
			}

		default:

		}
	}

	// 2. asKElected
		/*
		si (moi, monApt) ∈liste alors
		élu := i tel que apti= max(aptj)
		∀j dans liste
		// max(i) arbitraire à égalité
		envoie RESULTAT(élu,{moi})
		état := résultat
		sinon
		envoie ANNONCE( {(moi,monApt)} ∪liste)
		état := annonce
		fin si
		 */
	// 3. Annonce list
	// 4. Resultat (i, list)
		/*
		si moi ∈ liste alors état := non // fin
		sinon si état = résultat et élu <> i alors
		envoie ANNONCE({(moi,monApt)})
		état := annonce
		sinon si état = annonce alors
		élu := i
		envoie RESULTAT( élu, {moi} ∪ liste)
		état := résultat
		fin si
		 */
}

func (m *Manager) startElection(){

}

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
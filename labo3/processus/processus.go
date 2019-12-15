package processus

import (
	"math/rand"
	"prr-labo3/labo3/manager"
	"prr-labo3/labo3/network"
	"prr-labo3/labo3/task"
)

const (
	stampMax = 50
	stampMin = 1
)

type Processus struct {
	_N uint16
	id uint16
	network network.Network
	manager manager.Manager
	task task.Task
}

func (p *Processus)Init(id uint16, N uint16)  {
	p.id = id
	p._N = N
	p.network = network.Network{}
	p.manager = manager.Manager{}
	p.task = task.Task{}

	// Ensures everyone has a different seed
	rand.Seed(int64(id + N))
	aptitude := uint16(rand.Intn(stampMax - stampMin + 1) + stampMin)

	p.network.Init(p.id,p._N,&p.manager)
	p.manager.Init(p._N,p.id,aptitude,&p.network)
	p.task.Run(&p.manager,&p.network)
}

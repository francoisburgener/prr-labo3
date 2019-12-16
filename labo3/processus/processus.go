/*
 -----------------------------------------------------------------------------------
 Lab 		 : 03
 File    	 : processus.go
 Authors   	 : François Burgener - Tiago P. Quinteiro
 Date        : 10.12.19

 Goal        : Creates instances of network, manager and task
 -----------------------------------------------------------------------------------
*/
package processus

import (
	"log"
	"prr-labo3/labo3/manager"
	"prr-labo3/labo3/network"
	"prr-labo3/labo3/task"
)

type Processus struct {
	_N uint16
	id uint16
	network network.Network
	manager manager.Manager
	task task.Task
}

func (p *Processus)Init(id uint16, N uint16, aptitude uint16)  {
	p.id = id
	p._N = N
	p.network = network.Network{
		Debug: false,
	}
	p.manager = manager.Manager{}
	p.task = task.Task{}


	log.Println("Aptitude is ", aptitude)

	p.network.Init(p.id,p._N,&p.manager)
	p.manager.Init(p._N,p.id,aptitude,&p.network)
	p.task.Run(&p.manager,&p.network)
}

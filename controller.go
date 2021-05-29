package flow

import "github.com/morrocker/broadcast"

// type Gate struct{

// }
type Controller interface {
	Checkpoint() int
	Stop()
	Go()
	Exit(int)
}

type ControlTower struct {
	broadcaster *broadcast.Broadcaster
	allow       bool
	exitVal     int
}

func New() Controller {
	c := &ControlTower{
		broadcaster: broadcast.New(),
		allow:       true,
		exitVal:     0,
	}
	return c
}

func (c *ControlTower) Checkpoint() int {
	l := c.broadcaster.Listen()
	for {
		if c.allow {
			l.Close()
			return c.exitVal
		}
		<-l.C
	}
}

func (c *ControlTower) Stop() {
	c.allow = false
}

func (c *ControlTower) Go() {
	c.allow = true
	c.exitVal = 0
	c.broadcaster.Broadcast()
}

func (c *ControlTower) Exit(n int) {
	c.allow = true
	c.exitVal = n
	c.broadcaster.Broadcast()
}

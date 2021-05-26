package flow

import "github.com/morrocker/broadcast"

// type Gate struct{

// }
type Controller struct {
	broadcaster *broadcast.Broadcaster
	allow       bool
	exitVal     int
}

func New() *Controller {
	c := &Controller{
		broadcaster: broadcast.New(),
		allow:       true,
		exitVal:     0,
	}
	return c
}

func (c *Controller) Checkpoint() int {
	l := c.broadcaster.Listen()
	for {
		if c.allow {
			l.Close()
			return c.exitVal
		}
		<-l.C
	}
}

func (c *Controller) Stop() {
	c.allow = false
}

func (c *Controller) Go() {
	c.allow = true
	c.exitVal = 0
	c.broadcaster.Broadcast()
}

func (c *Controller) Exit(n int) {
	c.allow = true
	c.exitVal = n
	c.broadcaster.Broadcast()
}

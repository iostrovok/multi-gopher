package queues

/*
	Package supports the IChan interface and provides 3 simples realization of them.

	- standard GO channel FIFO, the fastest realization
	- priority queues
	- stack, LIFO
*/

import (
	"github.com/iostrovok/conveyor/faces"
	"github.com/iostrovok/conveyor/queues/priorityqueue"
	"github.com/iostrovok/conveyor/queues/stack"
	"github.com/iostrovok/conveyor/queues/std"
)

func New(lengthChannel int, chanType faces.ChanType) faces.IChan {
	switch chanType {
	case faces.ChanStdGo:
		return std.New(lengthChannel)
	case faces.ChanStack:
		return stack.New(lengthChannel)
	case faces.ChaPriorityQueue:
		return priorityqueue.New(lengthChannel)
	}

	return nil
}

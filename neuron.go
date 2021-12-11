package main

import "math"

type neuron struct {
	id         int
	sum        float64
	layer      int
	_activated bool //assigned by internal functions
}

func (n *neuron) reset() {
	n._activated = false
	n.sum = 0
}

func (n *neuron) activationFunction() {
	if !n._activated {
		n.sum = math.Max(0, n.sum) //Simple RELu function
		n._activated = true
	}

}

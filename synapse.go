package main

type synapse struct { //fancy for connection
	from   *neuron
	to     *neuron
	weight float64
}

func (s *synapse) process() {
	s.from.activationFunction()
	s.to.sum += (s.from.sum * s.weight)
}

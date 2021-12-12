package main

type synapse struct { //fancy for connection
	from   *neuron
	to     *neuron
	weight float64
}

func NewSynapseP(f *neuron, t *neuron, w float64) *synapse {
	return &synapse{
		f,
		t,
		w,
	}
}

func (s *synapse) process() {
	s.from.activationFunction()
	s.to.sum += (s.from.sum * s.weight)
}

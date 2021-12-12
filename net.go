package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type NeuralNet struct {
	inputs      []*neuron
	connections []*synapse
	neurons     [][]*neuron
	outputs     []*neuron
}

//NewNet Standard feed-forward network is generated
func NewNet(inputSize int, layerCount int, perLayer int, outputSize int) (NeuralNet, error) {
	rand.Seed(time.Now().UnixNano())
	if inputSize <= 0 || outputSize <= 0 {
		return NeuralNet{}, fmt.Errorf("Invalid input/output size, must be greater than zero!")
	}

	var in []*neuron
	for i := 0; i < inputSize; i++ {
		in = append(in, NewNeuronP(-1))
	}

	var connections []*synapse
	var structure [][]*neuron //only used to make life easier when adding all the connections

	//Create required Neurons
	nextId := 0
	for i := 0; i < layerCount; i++ {
		structure = append(structure, []*neuron{})
		for j := 0; j < perLayer; j++ {
			structure[i] = append(structure[i], NewNeuronP(nextId))
			nextId++
		}
	}

	//Connect each layer, n^3 complexity... room for optimization
	//TODO: Use Parallel computing to generate each layer
	for i := 0; i < layerCount; i++ {
		for l := 0; l < perLayer; l++ {
			neuron := structure[i][l]

			if i-1 >= 0 { //Connect to previous layer
				for j := 0; j < perLayer; j++ {
					connections = append(connections, NewSynapseP(structure[i-1][j], neuron, rand.Float64()/math.MaxFloat64))
				}
			} else { //Connect to inputs
				for j := 0; j < inputSize; j++ {
					connections = append(connections, NewSynapseP(in[j], neuron, rand.Float64()/math.MaxFloat64))
				}
			}
		}
	}

	//Finish with Outputs, and Connect them
	var out []*neuron
	for i := 0; i < outputSize; i++ {
		temp := NewNeuronP(-1)
		in = append(out, temp)
		for j := 0; j < perLayer; j++ {
			connections = append(connections, NewSynapseP(structure[len(structure)-1][j], temp, rand.Float64()/math.MaxFloat64))
		}
	}

	return NeuralNet{
		in,
		connections,
		structure,
		out,
	}, nil
}

//FeedForward Feeds the inputs forward
func (n *NeuralNet) FeedForward(inputs ...float64) error {
	if len(inputs) != len(n.inputs) {
		return fmt.Errorf("Input amount did not match network input size.")
	}

	//Load into inputs
	for i := 0; i < len(n.inputs); i++ {
		n.inputs[i].sum = inputs[i]
	}

	//Find each connection, and process
	for i := 0; i < len(n.neurons); i++ { //Process each layer sequentially (Don't do Multi-threading)
		layer := n.neurons[i]
		for _, neuron := range layer {
			neuron.reset()

			//TODO: Separate this into new threads
			for _, syn := range n.connections { //TODO: Total Connections Variable that way we can
				if syn.to == neuron {
					syn.process()
				}
			}
		}
	}

	//NOTE: This needs to looked at when we add a total connections variable
	maxConnections := len(n.neurons[len(n.neurons)-1])
	for _, out := range n.outputs {
		conAm := 0
		for _, syn := range n.connections {
			if syn.to == out {
				syn.process()
			}

			conAm++

			if conAm >= maxConnections {
				break
			}
		}
	}

	return nil
}

func (n *NeuralNet) GetSimplifiedOutputs() (outs []float64) {
	for _, o := range n.outputs {
		o.activationFunction()
		outs = append(outs, o.sum)
	}
	return
}

func (n *NeuralNet) ConsoleDisplay() {
	fmt.Print("Inputs:")
	for _, i := range n.inputs {
		fmt.Printf(" %f", i.sum)
	}
	fmt.Println()

	for i := 0; i < len(n.neurons); i++ {
		fmt.Printf("Layer %v:", i+1)
		layer := n.neurons[i]
		for _, neuron := range layer {
			fmt.Printf(" %f", neuron.sum)
		}
	}
	fmt.Println()

	fmt.Print("Outputs:")
	for _, i := range n.outputs {
		fmt.Printf(" %f", i.sum)
	}
	fmt.Println()
}
